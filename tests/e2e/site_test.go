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

func Test_AltaClient_Site(t *testing.T) {
	client := SetupTestAltaClient()

	t.Run("ListSites should return sites", func(t *testing.T) {
		sites, err := client.ListSites()
		require.NoError(t, err)
		require.NotEmpty(t, sites)
	})

	t.Run("Rename", func(t *testing.T) {
		t.Run("Rename by name", func(t *testing.T) {
			err := client.RenameSite("RenameSiteTest", "RenameSite2Test")
			require.NoError(t, err)

			sites, err := client.ListSites()
			require.NoError(t, err)
			found1 := false
			for _, site := range sites {
				if site.Name == "RenameSite2Test" {
					found1 = true
				}
			}
			require.True(t, found1)

			err = client.RenameSite("RenameSite2Test", "RenameSiteTest")
			require.NoError(t, err)

			sites, err = client.ListSites()
			require.NoError(t, err)
			found2 := false
			for _, site := range sites {
				if site.Name == "RenameSiteTest" {
					found2 = true
				}
			}
			require.True(t, found2)
		})

	})

	t.Run("GetSite should return site", func(t *testing.T) {
		sites, err := client.ListSites()
		require.NoError(t, err)
		require.NotEmpty(t, sites)

		var siteID string
		for _, site := range sites {
			if site.Name == "GetSiteTest" {
				siteID = site.ID
				break
			}
		}

		require.NotEmpty(t, siteID)
		site, err := client.GetSite(siteID)
		require.NoError(t, err)
		require.NotEmpty(t, site)
		require.Equal(t, "getuser", site.Username)
	})
}
