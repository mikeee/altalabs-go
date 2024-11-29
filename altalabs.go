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
	"context"
	"errors"
	"fmt"
	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"io"
	"log"
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
	userConfig Config
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

func (a *AuthClient) SignIn(config *Config) error {
	srp, err := cognitosrp.NewCognitoSRP(config.Username, config.Password, a.userPoolID, a.clientID, a.clientSecret)
	if err != nil {
		return fmt.Errorf("failed to create cognito srp: %w", err)
	}

	resp, err := a.cognito.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       &a.clientID,
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

		respAuth, err := a.cognito.RespondToAuthChallenge(context.TODO(), &cognitoidentityprovider.RespondToAuthChallengeInput{
			ChallengeName:      types.ChallengeNameTypePasswordVerifier,
			ClientId:           aws.String(srp.GetClientId()),
			ChallengeResponses: respPV,
		})
		if err != nil {
			return fmt.Errorf("failed to respond to auth challenge: %w", err)
		}

		a.auth = respAuth.AuthenticationResult
		a.userConfig = *config

		return nil

	default:
		return errors.New("Unhandled auth challenge received: " + string(resp.ChallengeName))
	}
}

func (a *AuthClient) RefreshAuth() error {
	if err := a.SignIn(&a.userConfig); err != nil {
		return fmt.Errorf("failed to refresh auth: %w", err)
	}

	return nil
}

func (a *AuthClient) GetIDToken() string {
	if a.auth != nil {
		return *a.auth.IdToken
	}

	return ""
}

type AltaClient struct {
	client     *http.Client
	authClient *AuthClient
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
		authClient: authClient,
	}, nil
}

func (a *AltaClient) request(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, API_BASE_URL+url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (a *AltaClient) getRequest(path string) (*http.Response, error) {
	req, err := a.request(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Token", a.authClient.GetIDToken())

	return a.client.Do(req)
}

func (a *AltaClient) ListSites() (Sites, error) {
	siteURL := "sites/list"

	resp, err := a.getRequest(siteURL)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("failed to close response body: %v", err)
		}
	}()

	var sites = make(Sites, 0)
	if err := sites.Unmarshal(resp.Body); err != nil {
		return nil, fmt.Errorf("failed to decode sites: %w", err)
	}

	return sites, nil
}
