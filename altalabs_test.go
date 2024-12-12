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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestConfig(t *testing.T) {
	configExample := &Config{
		Username: "username",
		Password: "password",
	}
	t.Run("Config should be valid", func(t *testing.T) {
		config := NewConfig().WithSRPAuth(configExample.Username, configExample.Password)

		assert.Equal(t, configExample, config)
	})
}

func TestAltaClient(t *testing.T) {
	testClient := AltaClient{
		client:     nil,
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
		req, err := testClient.request("GET", "https://manage.alta.inc/api/", nil)
		require.NoError(t, err)
		assert.Equal(t, testRequest.Method, req.Method)
	})
}
