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
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	configExample := &Config{
		Username: "username",
		Password: "password",
	}
	t.Run("Config should be valid", func(t *testing.T) {
		testConfig := NewConfig().WithSRPAuth(configExample.Username, configExample.Password)

		assert.Equal(t, configExample, testConfig)
	})
}

func TestAuthClient(t *testing.T) {
	testAuth := &AuthClient{
		authConfig: nil,
		userConfig: &Config{},
		cognito:    nil,
		auth: &types.AuthenticationResultType{
			AccessToken:       nil,
			ExpiresIn:         0,
			IdToken:           nil,
			NewDeviceMetadata: nil,
			RefreshToken:      nil,
			TokenType:         nil,
		},
	}

	t.Run("Test expiry retrieval", func(t *testing.T) {
		testAuth.auth.ExpiresIn = int32(time.Now().Unix()) - 10 // Insert a valid expiry
		assert.Equal(t, testAuth.auth.ExpiresIn, testAuth.GetExpiry())
	})
}

func TestAltaClient(t *testing.T) {
	testClient := AltaClient{
		client: &http.Client{
			Transport:     http.DefaultTransport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		AuthClient: nil,
	}

	t.Run("Test request builder", func(t *testing.T) {
		testURL, err := url.Parse("https://manage.alta.inc/api/")
		require.NoError(t, err, "Failed to parse URL in test request builder")
		testRequest := http.Request{
			Method: "GET",
			URL:    testURL,
			Body:   nil,
		}

		testString := "abc123"
		testIDToken := "id_token"

		cognitoConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		require.NoError(t, err)

		testClient.AuthClient = &AuthClient{
			authConfig: &authConfig{
				userPoolID:   "abc_123",
				clientID:     testString,
				clientSecret: &testString,
			},
			userConfig: &Config{
				Username: "",
				Password: "",
			},
			cognito: cognitoidentityprovider.NewFromConfig(cognitoConfig),
			auth: &types.AuthenticationResultType{
				AccessToken: &testString,
				ExpiresIn:   int32(time.Now().Unix()) + 10, // Insert a valid expiry
				IdToken:     &testIDToken,
				NewDeviceMetadata: &types.NewDeviceMetadataType{
					DeviceGroupKey: &testString,
					DeviceKey:      &testString,
				},
				RefreshToken: &testString,
				TokenType:    &testString,
			},
		}

		t.Run("GET-style request", func(t *testing.T) {

			req, err := testClient.request("GET", "https://manage.alta.inc/api/", nil)
			require.NoError(t, err)
			assert.Equal(t, testRequest.Method, req.Method)
			assert.Equal(t, testRequest.URL.Host, req.URL.Host)
			assert.Equal(t, testRequest.Body, req.Body)
		})

		t.Run("POST-style request", func(t *testing.T) {
			serialisedBodyBytes := []byte(`{"key":"value"}`)
			postBodyBytes := []byte(`{"key":"value","token":"` + testIDToken + `"}`)

			req, err := testClient.request("POST", "https://manage.alta.inc/api/", serialisedBodyBytes)
			require.NoError(t, err)
			assert.Equal(t, "POST", req.Method)
			assert.Equal(t, testRequest.URL.Host, req.URL.Host)

			body, errReadBody := io.ReadAll(req.Body)
			require.NoError(t, errReadBody)
			errCloseBody := req.Body.Close()
			require.NoError(t, errCloseBody)
			assert.Equal(t, postBodyBytes, body, "POST-style request body is not as expected")
		})
	})

	t.Run("Test checkToken", func(t *testing.T) {
		t.Run("unexpired", func(t *testing.T) {
			testClient.AuthClient = &AuthClient{
				authConfig: &authConfig{
					userPoolID:   "",
					clientID:     "",
					clientSecret: nil,
				},
				userConfig: &Config{},
				cognito:    nil,
				auth: &types.AuthenticationResultType{
					AccessToken:       nil,
					ExpiresIn:         int32(time.Now().Unix()) + 10, // Insert a valid expiry
					IdToken:           nil,
					NewDeviceMetadata: nil,
					RefreshToken:      nil,
					TokenType:         nil,
				},
			}
			err := testClient.checkToken()
			require.NoError(t, err)
		})

		t.Run("expired", func(t *testing.T) {

			testClient.AuthClient = &AuthClient{
				authConfig: nil,
				userConfig: &Config{},
				cognito:    nil,
				auth: &types.AuthenticationResultType{
					AccessToken:       nil,
					ExpiresIn:         int32(time.Now().Unix()) - 10, // Insert a valid expiry
					IdToken:           nil,
					NewDeviceMetadata: nil,
					RefreshToken:      nil,
					TokenType:         nil,
				},
			}
			err := testClient.checkToken()
			require.Error(t, err)
		})
	})
}
