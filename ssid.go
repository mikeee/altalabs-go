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

import "fmt"

type SSIDList struct {
	SSIDs []SSID `json:"ssids"`
}

type SSID struct {
	Ssid   string   `json:"ssid"`
	Sites  []string `json:"sites"`
	Emails []string `json:"emails"`
	ID     string   `json:"id"`
	Config struct {
		Acl       string   `json:"acl"`
		Type      string   `json:"type"`
		Wpa3      string   `json:"wpa3"`
		Bands     string   `json:"bands"`
		Dtim2     int      `json:"dtim2"`
		Dtim5     int      `json:"dtim5"`
		Notes     string   `json:"notes"`
		Colors    []string `json:"colors"`
		AclList   string   `json:"aclList"`
		Network   int      `json:"network"`
		RadiusIP  string   `json:"radiusIp"`
		Schedule  string   `json:"schedule"`
		Security  string   `json:"security"`
		Passwords []struct {
			Vlan          int    `json:"vlan,omitempty"`
			DlRate        int    `json:"dlRate,omitempty"`
			Locked        bool   `json:"locked,omitempty"`
			UlRate        int    `json:"ulRate,omitempty"`
			Network       string `json:"network"`
			Password      string `json:"password"`
			IgnoreSched   bool   `json:"ignoreSched,omitempty"`
			IgnoreFilter  bool   `json:"ignoreFilter,omitempty"`
			IgnoreHotspot bool   `json:"ignoreHotspot,omitempty"`
		} `json:"passwords"`
		HotspotExt        string        `json:"hotspotExt"`
		HotspotType       string        `json:"hotspotType"`
		HotspotTerms      string        `json:"hotspotTerms"`
		HotspotTitle      string        `json:"hotspotTitle"`
		RadiusSecret      string        `json:"radiusSecret"`
		HotspotFinish     string        `json:"hotspotFinish"`
		HotspotSecret     string        `json:"hotspotSecret"`
		PowerSettings     string        `json:"powerSettings"`
		HotspotExtAuth    string        `json:"hotspotExtAuth"`
		RadiusAcctPort    int           `json:"radiusAcctPort"`
		RadiusAuthPort    int           `json:"radiusAuthPort"`
		ScheduleBlocks    []interface{} `json:"scheduleBlocks"` // TODO: Figure out what this is
		HotspotPassword   string        `json:"hotspotPassword"`
		HotspotTermsTitle string        `json:"hotspotTermsTitle"`
	} `json:"config"`
	Ftkey            string      `json:"ftkey"`
	NotifiedTemplate interface{} `json:"notifiedTemplate"`
}

func (a *AltaClient) ListSSID() (SSIDList, error) {
	URL := "wifi/ssid/list"

	var ssidList SSIDList
	err := a.getRequest(URL, nil, &ssidList)

	if err != nil {
		return SSIDList{}, err
	}
	return ssidList, nil
}

type NewSSIDRequest struct {
	Config struct {
		ID        string `json:"id"`
		Ssid      string `json:"ssid"`
		Notes     string `json:"notes"`
		Security  string `json:"security"`
		Bands     string `json:"bands"`
		Passwords []struct {
			Network       string `json:"network"`
			Password      string `json:"password"`
			Vlan          int    `json:"vlan"`
			DlRate        int    `json:"dlRate"`
			UlRate        int    `json:"ulRate"`
			IgnoreHotspot bool   `json:"ignoreHotspot"`
			IgnoreSched   bool   `json:"ignoreSched"`
			IgnoreFilter  bool   `json:"ignoreFilter"`
			Locked        bool   `json:"locked"`
		} `json:"passwords"`
		Network           int           `json:"network"`
		Wpa3              string        `json:"wpa3"`
		Dtim2             int           `json:"dtim2"`
		Dtim5             int           `json:"dtim5"`
		Schedule          string        `json:"schedule"`
		ScheduleBlocks    []interface{} `json:"scheduleBlocks"` // TODO: Figure out what this is
		Acl               string        `json:"acl"`
		AclList           string        `json:"aclList"`
		Type              string        `json:"type"`
		Colors            []string      `json:"colors"`
		RadiusIP          string        `json:"radiusIp"`
		RadiusSecret      string        `json:"radiusSecret"`
		RadiusAuthPort    int           `json:"radiusAuthPort"`
		RadiusAcctPort    int           `json:"radiusAcctPort"`
		HotspotTitle      string        `json:"hotspotTitle"`
		HotspotPassword   string        `json:"hotspotPassword"`
		HotspotTerms      string        `json:"hotspotTerms"`
		HotspotTermsTitle string        `json:"hotspotTermsTitle"`
		HotspotFinish     string        `json:"hotspotFinish"`
		HotspotExt        string        `json:"hotspotExt"`
		HotspotSecret     string        `json:"hotspotSecret"`
		HotspotExtAuth    string        `json:"hotspotExtAuth"`
		PowerSettings     string        `json:"powerSettings"`
		Sites             []string      `json:"sites"` // ID of the site
	} `json:"config"`
}

type NewSSIDResponse struct {
	ID string `json:"id"`
}

type GetSSIDRequest struct {
	ID string
}

type GetSSIDResponse struct {
	Ssid   string   `json:"ssid"`
	Sites  []string `json:"sites"`
	Emails []string `json:"emails"`
	ID     string   `json:"id"`
	Config struct {
		ACL       string   `json:"acl"`
		Type      string   `json:"type"`
		Wpa3      string   `json:"wpa3"`
		Bands     string   `json:"bands"`
		Dtim2     int      `json:"dtim2"`
		Dtim5     int      `json:"dtim5"`
		Notes     string   `json:"notes"`
		Colors    []string `json:"colors"`
		ACLList   string   `json:"aclList"`
		Network   int      `json:"network"`
		RadiusIP  string   `json:"radiusIp"`
		Schedule  string   `json:"schedule"`
		Security  string   `json:"security"`
		Passwords []struct {
			Network  string `json:"network"`
			Password string `json:"password"`
		} `json:"passwords"`
		HotspotExt        string `json:"hotspotExt"`
		HotspotType       string `json:"hotspotType"`
		HotspotTerms      string `json:"hotspotTerms"`
		HotspotTitle      string `json:"hotspotTitle"`
		RadiusSecret      string `json:"radiusSecret"`
		HotspotFinish     string `json:"hotspotFinish"`
		HotspotSecret     string `json:"hotspotSecret"`
		PowerSettings     string `json:"powerSettings"`
		HotspotExtAuth    string `json:"hotspotExtAuth"`
		RadiusAcctPort    int    `json:"radiusAcctPort"`
		RadiusAuthPort    int    `json:"radiusAuthPort"`
		ScheduleBlocks    []any  `json:"scheduleBlocks"`
		HotspotPassword   string `json:"hotspotPassword"`
		HotspotTermsTitle string `json:"hotspotTermsTitle"`
	} `json:"config"`
	Ftkey            string `json:"ftkey"`
	NotifiedTemplate any    `json:"notifiedTemplate"`
}

type EditSSIDRequest struct {
	Config struct {
		ID        string `json:"id,omitempty"`
		Ssid      string `json:"ssid,omitempty"`
		Notes     string `json:"notes,omitempty"`
		Security  string `json:"security,omitempty"`
		Bands     string `json:"bands,omitempty"`
		Passwords []struct {
			Network  string `json:"network,omitempty"`
			Password string `json:"password,omitempty"`
		} `json:"passwords,omitempty"`
		Network           int      `json:"network,omitempty"`
		Wpa3              string   `json:"wpa3,omitempty"`
		Dtim2             int      `json:"dtim2,omitempty"`
		Dtim5             int      `json:"dtim5,omitempty"`
		Schedule          string   `json:"schedule,omitempty"`
		ScheduleBlocks    []any    `json:"scheduleBlocks,omitempty"`
		ACL               string   `json:"acl,omitempty"`
		ACLList           string   `json:"aclList,omitempty"`
		Type              string   `json:"type,omitempty"`
		Colors            []string `json:"colors,omitempty"`
		RadiusIP          string   `json:"radiusIp,omitempty"`
		RadiusSecret      string   `json:"radiusSecret,omitempty"`
		RadiusAuthPort    int      `json:"radiusAuthPort,omitempty"`
		RadiusAcctPort    int      `json:"radiusAcctPort,omitempty"`
		HotspotType       string   `json:"hotspotType,omitempty"`
		HotspotTitle      string   `json:"hotspotTitle,omitempty"`
		HotspotPassword   string   `json:"hotspotPassword,omitempty"`
		HotspotTerms      string   `json:"hotspotTerms,omitempty"`
		HotspotTermsTitle string   `json:"hotspotTermsTitle,omitempty"`
		HotspotFinish     string   `json:"hotspotFinish,omitempty"`
		HotspotExt        string   `json:"hotspotExt,omitempty"`
		HotspotSecret     string   `json:"hotspotSecret,omitempty"`
		HotspotExtAuth    string   `json:"hotspotExtAuth,omitempty"`
		PowerSettings     string   `json:"powerSettings,omitempty"`
		Sites             []string `json:"sites,omitempty"`
	} `json:"config,omitempty"`
}

func (a *AltaClient) GetSSID(id string) (*GetSSIDResponse, error) {
	URL := "wifi/ssid"
	var req = GetSSIDRequest{ID: id}

	var resp GetSSIDResponse

	if err := a.getRequest(URL, req, &resp); err != nil {
		return nil, fmt.Errorf("failed to get SSID: %w", err)
	}

	return &resp, nil
}

func (a *AltaClient) AddSSID(req NewSSIDRequest) (*string, error) {
	URL := "wifi/ssid"

	var resp NewSSIDResponse

	if err := a.postRequest(URL, req, &resp); err != nil {
		return nil, fmt.Errorf("failed to add SSID: %w", err)
	}

	return &resp.ID, nil
}

func (a *AltaClient) EditSSID(req EditSSIDRequest) error {
	URL := "wifi/ssid"

	if err := a.postRequest(URL, req, nil); err != nil {
		return fmt.Errorf("failed to edit SSID: %w", err)
	}

	return nil
}

// TODO: Consider xnet ssid creation

func (a *AltaClient) DeleteSSID(id string) error {
	URL := "wifi/ssid/delete"

	req := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	if err := a.postRequest(URL, req, nil); err != nil {
		return fmt.Errorf("failed to delete SSID: %w", err)
	}

	return nil
}
