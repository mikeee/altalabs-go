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
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_AltaClient_Sites(t *testing.T) {
	client := SetupTestAltaClient()

	t.Run("ListSites should return sites", func(t *testing.T) {
		sites, err := client.ListSites()
		require.NoError(t, err)
		require.NotEmpty(t, sites)
	})
}
