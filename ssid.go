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

func (a *AltaClient) ListSSIDs() (SSIDList, error) {
	URL := "wifi/ssid/list"

	var ssidList SSIDList
	err := a.getRequest(URL, &ssidList)

	if err != nil {
		return SSIDList{}, err
	}
	return ssidList, nil
}
