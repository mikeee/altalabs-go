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
	"errors"
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

type newSiteRequest struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Tz   string `json:"tz"`
}

type newSiteResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateSiteOption func(*newSiteRequest)

func WithSiteIcon(icon string) CreateSiteOption {
	return func(req *newSiteRequest) {
		req.Icon = icon
	}
}

func WithSiteTz(tz string) CreateSiteOption {
	return func(req *newSiteRequest) {
		req.Tz = tz
	}
}

func (a *AltaClient) CreateSite(name string, opts ...CreateSiteOption) (newSiteResponse, error) {
	newSite := newSiteRequest{
		Name: name,
	}

	for _, opt := range opts {
		opt(&newSite)
	}

	var resp newSiteResponse
	err := a.postRequest("sites/new", newSite, &resp)
	if err != nil {
		return newSiteResponse{}, err
	}
	return resp, nil
}

type GetSiteRequest struct {
	Id string
}

func (a *AltaClient) GetSite(siteID string) (*Site, error) {
	reqParams := GetSiteRequest{
		Id: siteID,
	}
	var site Site
	err := a.getRequest("site", reqParams, &site)
	if err != nil {
		return nil, err
	}
	return &site, nil
}

type renameSiteRequest struct {
	SiteID string `json:"siteid"`
	Name   string `json:"name"`
}

func (a *AltaClient) RenameSite(old, new string) error {
	sites, err := a.ListSites()
	if err != nil {
		return err
	}

	// find the site id from sites
	var siteID string
	for _, site := range sites {
		if site.Name == old {
			siteID = site.ID
			break
		}
	}

	if siteID == "" {
		return errors.New("site not found")
	}

	return a.RenameSiteByID(siteID, new)
}

func (a *AltaClient) RenameSiteByID(siteID, name string) error {
	req := renameSiteRequest{
		SiteID: siteID,
		Name:   name,
	}

	if err := a.postRequest("sites/rename", req, nil); err != nil {
		return err
	}

	return nil
}

func (a *AltaClient) UpdateSite(site Site) error {
	if err := a.postRequest("sites/update", site, nil); err != nil {
		return err
	}
	return nil
}
