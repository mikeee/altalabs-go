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
	"io"
)

type Site struct {
	ID          string   `json:"id"`
	Tz          string   `json:"tz"`
	SSHKeys     []string `json:"sshKeys"`
	Iappkey     string   `json:"iappkey"`
	Meshid      string   `json:"meshid"`
	Meshpw      string   `json:"meshpw"`
	BlockedApps struct {
		List       []string `json:"list"`
		Selections []string `json:"selections"`
	} `json:"blockedApps"`
	Leds              string   `json:"leds"`
	Viewers           any      `json:"viewers"`
	Vlans             any      `json:"vlans"`
	PortColors        any      `json:"portColors"`
	Radii             any      `json:"radii"`
	Update            bool     `json:"update"`
	SwitchLeds        string   `json:"switchLeds"`
	SyslogHost        string   `json:"syslogHost"`
	Backnet           any      `json:"backnet"`
	DpiEngine         bool     `json:"dpiEngine"`
	StpMode           any      `json:"stpMode"`
	DisconnectTimeout any      `json:"disconnectTimeout"`
	Community         string   `json:"community"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
	Contact           string   `json:"contact"`
	AutoDfs           bool     `json:"autoDfs"`
	Rchans2           []int    `json:"rchans2"`
	Rchans5           []int    `json:"rchans5"`
	Wans              any      `json:"wans"`
	Firewall          Firewall `json:"firewall"`
	DhcpGuard         string   `json:"dhcpGuard"`
	DhcpMacList       any      `json:"dhcpMacList"`
	Width2            any      `json:"width2"`
	Width5            any      `json:"width5"`
	AllowNewUsers     bool     `json:"allowNewUsers"`
}

func (s *Site) Unmarshal(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(s)
}
