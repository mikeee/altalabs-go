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

package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateTokenPair(t *testing.T) {
	token := "test"
	tokenExpected := `"token":"test"`
	t.Run("GenerateTokenPair should return a token", func(t *testing.T) {
		tokenPair, err := GenerateTokenPair(token)
		require.NoError(t, err)
		assert.Equal(t, tokenExpected, tokenPair)
	})
}

func TestAppendTokenToJSONBody(t *testing.T) {
	body := []byte(`{"test":"test"}`)
	token := `"token":"token"`
	bodyExpected := []byte(`{"test":"test","token":"token"}`)

	t.Run("AppendTokenToJSONBody should append a token to a JSON body", func(t *testing.T) {
		bodyWithToken, err := AppendTokenToJSONBody(body, token)
		require.NoError(t, err)
		assert.Equal(t, bodyExpected, bodyWithToken)
	})
}
