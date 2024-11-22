package altalabs

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSite(t *testing.T) {
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
	var sites []Site

	if err := json.Unmarshal([]byte(jsonData), &sites); err != nil {
		require.NoError(t, err, "error unmarshalling json")
	}
	assert.Equal(t, "office", sites[1].Name)

}
