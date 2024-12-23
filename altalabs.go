/*
Copyright 2024 Mike Nguyen (mikeee) <hey@mike.ee>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package altalabs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/mikeee/altalabs-go/util"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"
)

const (
	API_BASE_URL = "https://manage.alta.inc/api/"
)

type Method int

const (
	ALTA_CLIENT_ID = "24bk8l088t5bf31nuceoqb503q"

	COGNITO_REGION       = "us-east-1"
	COGNITO_USER_POOL_ID = "4QbA7N3Uy"
)

type Config struct {
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithSRPAuth(username, password string) *Config {
	c.Username = username
	c.Password = password
	return c
}

type authConfig struct {
	userPoolID   string
	clientID     string
	clientSecret *string
}

type AuthClient struct {
	*authConfig
	userConfig *Config
	cognito    *cognitoidentityprovider.Client
	auth       *types.AuthenticationResultType
}

func NewAuthClient(region string) (*AuthClient, error) {
	authConfig := authConfig{
		userPoolID:   COGNITO_REGION + "_" + COGNITO_USER_POOL_ID,
		clientID:     ALTA_CLIENT_ID,
		clientSecret: nil,
	}

	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load aws config: %w", err)
	}

	return &AuthClient{
		authConfig: &authConfig,
		cognito:    cognitoidentityprovider.NewFromConfig(awsConfig),
	}, nil
}

func (auth *AuthClient) SignIn(config *Config) error {
	if config == nil {
		return errors.New("config is nil")
	}
	srp, err := cognitosrp.NewCognitoSRP(config.Username, config.Password, auth.userPoolID, auth.clientID, auth.clientSecret)
	if err != nil {
		return fmt.Errorf("failed to create cognito srp: %w", err)
	}

	resp, err := auth.cognito.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       &auth.clientID,
		AuthParameters: srp.GetAuthParams(),
	})
	if err != nil {
		return fmt.Errorf("failed to initiate auth: %w", err)
	}

	switch resp.ChallengeName {
	case types.ChallengeNameTypePasswordVerifier:
		respPV, err := srp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())
		if err != nil {
			return fmt.Errorf("failed to verify password: %w", err)
		}

		respAuth, err := auth.cognito.RespondToAuthChallenge(context.TODO(), &cognitoidentityprovider.RespondToAuthChallengeInput{
			ChallengeName:      types.ChallengeNameTypePasswordVerifier,
			ClientId:           aws.String(srp.GetClientId()),
			ChallengeResponses: respPV,
		})
		if err != nil {
			return fmt.Errorf("failed to respond to auth challenge: %w", err)
		}

		auth.auth = respAuth.AuthenticationResult
		auth.userConfig = config

		return nil

	default:
		return errors.New("Unhandled auth challenge received: " + string(resp.ChallengeName))
	}
}

func (auth *AuthClient) RefreshAuth() error {
	if auth.userConfig == nil {
		return errors.New("user config is nil")
	}
	if err := auth.SignIn(auth.userConfig); err != nil {
		return fmt.Errorf("failed to refresh auth: %w", err)
	}

	return nil
}

func (auth *AuthClient) GetIDToken() string {
	if auth.auth != nil {
		return *auth.auth.IdToken
	}

	return ""
}

func (auth *AuthClient) GetExpiry() int32 {
	if auth.auth != nil {
		return auth.auth.ExpiresIn
	}

	return 0
}

type AltaClient struct {
	client     *http.Client
	AuthClient *AuthClient
}

func NewAltaClient(username, password string) (*AltaClient, error) {
	authClient, err := NewAuthClient(COGNITO_REGION)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth client: %w", err)
	}

	clientConfig := NewConfig().WithSRPAuth(username, password)

	if err := authClient.SignIn(clientConfig); err != nil {
		return nil, fmt.Errorf("failed to sign in: %w", err)
	}

	return &AltaClient{
		client:     &http.Client{},
		AuthClient: authClient,
	}, nil
}

var (
	ErrorAuthExpired = errors.New("auth token expired")
)

func (a *AltaClient) checkToken() error {
	// Check and renew tokens expired or within 5 seconds
	// TODO: This is a bad idea, refactor this
	if a.AuthClient.GetExpiry() <= int32(time.Now().Unix())+5 {
		return ErrorAuthExpired
	}
	return nil
}

func (a *AltaClient) request(method, url string, body []byte) (*http.Request, error) {
	if err := a.checkToken(); errors.Is(err, ErrorAuthExpired) {
		slog.Info("Refreshing auth token")
		if err := a.AuthClient.RefreshAuth(); err != nil {
			slog.Error("Failed to refresh auth token", slog.String("error", err.Error()))
		}
	}

	var reqBodyStream io.Reader
	var reqBody []byte

	if body != nil {
		// Append token to the payload/body
		// TODO: This is a bad idea, refactor this
		tokenPair, err := util.GenerateTokenPair(a.AuthClient.GetIDToken())
		if err != nil {
			return nil, fmt.Errorf("failed to generate token pair: %w", err)
		}

		reqBody, err = util.AppendTokenToJSONBody(body, tokenPair)
		if err != nil {
			return nil, fmt.Errorf("failed to append token to json body: %w", err)
		}

		reqBodyStream = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, API_BASE_URL+url, reqBodyStream)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (a *AltaClient) getRequest(path string, params, dest interface{}) error {
	if params != nil {
		path += "?" + util.StructToParams(params)
	}

	req, err := a.request(http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Token", a.AuthClient.GetIDToken())

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	if dest == nil {
		return nil
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("failed to close response body: %v", err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func (a *AltaClient) postRequest(path string, payload interface{}, dest interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := a.request(http.MethodPost, path, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	if dest == nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("failed to close response body: %v", err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
