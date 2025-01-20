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

import "errors"

type Firewall struct {
	Nat struct {
		Rules []any `json:"rules"`
	} `json:"nat"`
	Firewall struct {
		Rules []struct {
			ID          string   `json:"id"`
			Action      string   `json:"action"`
			ZoneIn      string   `json:"zoneIn"`
			IcmpType    []string `json:"icmpType,omitempty"`
			Protocol    []string `json:"protocol,omitempty"`
			IPVersion   string   `json:"ipVersion,omitempty"`
			Description string   `json:"description,omitempty"`
			Destination struct {
				Port any `json:"port"` // FIXME: The port can either be a string or int??
			} `json:"destination,omitempty"`
			Source struct {
				Address string `json:"address"`
			} `json:"source,omitempty"`
			Limit string `json:"limit,omitempty"`
		} `json:"rules"`
	} `json:"firewall"`
}

func (a *AltaClient) GetFirewall(siteID string) (*Firewall, error) {
	site, err := a.GetSite(siteID)
	if err != nil {
		return nil, err
	}

	return &site.Firewall, nil
}

func (a *AltaClient) UpdateFirewall() error {
	return errors.New("not implemented")
}

func (a *AltaClient) AddFirewallRule() error {
	return errors.New("not implemented")
}

func (a *AltaClient) DeleteFirewall() error {
	// Clear/empty the firewall rules
	return errors.New("not implemented")
}
