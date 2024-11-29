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
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Testsite(t *testing.T) {
	// Parse json into a site
	jsonData := `
[
    {
        "id": "E6WIX3G8vYq4g7d5",
        "name": "mikeee",
        "icon": "",
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    },
    {
        "id": "WKLFJK0F6hQ9q7b2",
        "name": "office",
        "icon": null,
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    },
    {
        "id": "YEUVWRG8hwx8d2l6",
        "name": "home",
        "icon": "",
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    }
]`
	t.Run("Site should be valid", func(t *testing.T) {
		var sites []site

		if err := json.Unmarshal([]byte(jsonData), &sites); err != nil {
			require.NoError(t, err, "error unmarshalling json")
		}
		assert.Equal(t, "office", sites[1].Name)
	})
}

func TestSites(t *testing.T) {
	// Parse json into a site
	jsonData := `
[
    {
        "id": "E6WIX3G8vYq4g7d5",
        "name": "mikeee",
        "icon": "",
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    },
    {
        "id": "WKLFJK0F6hQ9q7b2",
        "name": "office",
        "icon": null,
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    },
    {
        "id": "YEUVWRG8hwx8d2l6",
        "name": "home",
        "icon": "",
        "devices": [],
        "online": 0,
        "emails": [
            "altalabscloudtestemail@mike.ee"
        ],
        "perms": {
            "altalabscloudtestemail@mike.ee": {
                "admin": true,
                "allPasswords": true,
                "unlockedPasswords": true
            }
        }
    }
]`
	t.Run("Sites should be valid", func(t *testing.T) {
		var sites Sites

		err := sites.UnmarshalJSON(strings.NewReader(jsonData))
		require.NoError(t, err, "error unmarshalling json")
		assert.Equal(t, "office", sites[1].Name)
	})
}
