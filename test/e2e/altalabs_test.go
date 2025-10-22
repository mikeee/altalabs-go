//go:build e2e

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

package e2e

import (
	"github.com/mikeee/altalabs-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func Test_AuthClient(t *testing.T) {
	config := altalabs.NewConfig().WithSRPAuth(os.Getenv("SDK_ALTA_USER"), os.Getenv("SDK_ALTA_PASS"))
	t.Run("Config should be valid", func(t *testing.T) {
		configExample := &altalabs.Config{
			Username: os.Getenv("SDK_ALTA_USER"),
			Password: os.Getenv("SDK_ALTA_PASS"),
		}
		assert.Equal(t, configExample, config)
	})
	authClient, err := altalabs.NewAuthClient(altalabs.COGNITO_REGION)
	t.Run("Client should be valid", func(t *testing.T) {
		require.NoError(t, err)
		assert.NotNil(t, authClient)
	})

	t.Run("SignIn should be successful", func(t *testing.T) {
		err = authClient.SignIn(config)
		require.NoError(t, err)
	})

	t.Run("RefreshAuth should be successful", func(t *testing.T) {
		err = authClient.RefreshAuth()
		require.NoError(t, err)
	})

	t.Run("ID Token should be populated", func(t *testing.T) {
		assert.NotEmpty(t, authClient.GetIDToken())
	})
}

func Test_AltaClient(t *testing.T) {
	client, err := altalabs.NewAltaClient(os.Getenv("SDK_ALTA_USER"), os.Getenv("SDK_ALTA_PASS"))

	t.Run("Client should be valid and be successful with a ListSites request", func(t *testing.T) {
		require.NoError(t, err)
		assert.NotNil(t, client)

		sites, err := client.ListSites()
		require.NoError(t, err)
		assert.NotEmpty(t, sites)
	})

	t.Run("Refresh token should be successful with a list sites request", func(t *testing.T) {
		err = client.AuthClient.RefreshAuth()
		require.NoError(t, err)

		sites, err := client.ListSites()
		require.NoError(t, err)
		assert.NotEmpty(t, sites)
	})
}

func SetupTestAltaClient() *altalabs.AltaClient {
	client, err := altalabs.NewAltaClient(os.Getenv("SDK_ALTA_USER"), os.Getenv("SDK_ALTA_PASS"))
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}
	return client
}
