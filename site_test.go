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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSite(t *testing.T) {
	jsonData := `{
    "id": "vjnke435Temrr",
    "tz": "Europe/London",
    "sshKeys": [
        "ecdsa-sha2-nistp521 KEY== altalabstesting@mike.ee"
    ],
    "iappkey": "n2kj43nk24n3j4n2k3n4k2",
    "meshid": "REWREW2jk3n4jk2n",
    "meshpw": "rlwerWERj234j",
    "blockedApps": {
        "list": [
            "adobeconnect"
        ],
        "selections": [
            "adobeconnect"
        ]
    },
    "leds": "blue",
    "viewers": null,
    "vlans": null,
    "portColors": null,
    "radii": null,
    "update": true,
    "switchLeds": "#ff0000",
    "syslogHost": "192.168.0.1:80",
    "backnet": null,
    "dpiEngine": true,
    "stpMode": null,
    "disconnectTimeout": null,
    "community": "communitystring",
    "username": "usernamestring",
    "password": "passwordstring",
    "contact": "contact",
    "autoDfs": true,
    "rchans2": [
        2,
        3,
        4,
        5,
        7,
        8,
        9,
        10,
        12,
        13,
        14
    ],
    "rchans5": [
        40
    ],
    "wans": null,
    "firewall": {
        "nat": {
            "rules": []
        },
        "firewall": {
            "rules": [
                {
                    "id": "DSfsdf",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "icmpType": [
                        "echo-request"
                    ],
                    "protocol": [
                        "icmp"
                    ],
                    "ipVersion": "ipv4",
                    "description": "Allow Ping"
                },
                {
                    "id": "I034sx",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "protocol": [
                        "udp"
                    ],
                    "ipVersion": "ipv4",
                    "description": "Allow DHCP renewals",
                    "destination": {
                        "port": "68"
                    }
                },
                {
                    "id": "WobfdsM",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "protocol": [
                        "igmp"
                    ],
                    "ipVersion": "ipv4",
                    "description": "Allow IGMP"
                },
                {
                    "id": "Vsdfre",
                    "action": "ACCEPT",
                    "source": {
                        "address": "fc00::/6"
                    },
                    "zoneIn": "wan",
                    "protocol": [
                        "udp"
                    ],
                    "ipVersion": "ipv6",
                    "description": "Allow DHCPv6",
                    "destination": {
                        "port": 546,
                        "address": "fc00::/6"
                    }
                },
                {
                    "id": "kshDCk",
                    "action": "ACCEPT",
                    "source": {
                        "address": "fe80::/10"
                    },
                    "zoneIn": "wan",
                    "icmpType": [
                        "130/0",
                        "131/0",
                        "132/0",
                        "143/0"
                    ],
                    "protocol": [
                        "icmp"
                    ],
                    "ipVersion": "ipv6",
                    "description": "Allow MLD"
                },
                {
                    "id": "qSkeHW",
                    "limit": "1000/sec",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "icmpType": [
                        "echo-request",
                        "echo-reply",
                        "destination-unreachable",
                        "packet-too-big",
                        "time-exceeded",
                        "bad-header",
                        "unknown-header-type",
                        "router-solicitation",
                        "neighbour-solicitation",
                        "router-advertisement",
                        "neighbour-advertisement"
                    ],
                    "protocol": [
                        "icmp"
                    ],
                    "ipVersion": "ipv6",
                    "description": "Allow ICMPv6 input"
                },
                {
                    "id": "91m23E",
                    "limit": "1000/sec",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "icmpType": [
                        "echo-request",
                        "echo-reply",
                        "destination-unreachable",
                        "packet-too-big",
                        "time-exceeded",
                        "bad-header",
                        "unknown-header-type"
                    ],
                    "protocol": [
                        "icmp"
                    ],
                    "ipVersion": "ipv6",
                    "description": "Allow ICMPv6 forward"
                },
                {
                    "id": "XuCDgB",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "protocol": [
                        "udp"
                    ],
                    "description": "IPsec IKE",
                    "destination": {
                        "port": "500"
                    }
                },
                {
                    "id": "zf23YS",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "protocol": [
                        "udp"
                    ],
                    "description": "IPsec NAT-T",
                    "destination": {
                        "port": "4500"
                    }
                },
                {
                    "id": "78DESC",
                    "action": "ACCEPT",
                    "zoneIn": "wan",
                    "protocol": [
                        "esp"
                    ],
                    "description": "IPsec ESP"
                },
                {
                    "id": "FDEXsl",
                    "action": "ACCEPT",
                    "source": {
                        "address": "1.1.1.1"
                    },
                    "zoneIn": "wan"
                }
            ]
        }
    },
    "dhcpGuard": "auto",
    "dhcpMacList": null,
    "width2": null,
    "width5": null,
    "allowNewUsers": true
}`
	t.Run("Test site", func(t *testing.T) {
		var site Site
		err := site.UnmarshalJSON(strings.NewReader(jsonData))
		require.NoError(t, err, "error unmarshalling detailed site information")
		assert.Equal(t, "auto", site.DhcpGuard)
	})
}
