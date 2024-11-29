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

type Sites []site

type site struct {
	ID      string              `json:"id"`
	Name    string              `json:"name"`
	Icon    *string             `json:"icon"`
	Devices []interface{}       `json:"devices"` // TODO: implement devices struct
	Online  int                 `json:"online"`
	Emails  []string            `json:"emails"`
	Perms   map[string]sitePerm `json:"perms"`
}

type sitePerm struct {
	Admin             bool `json:"admin"`
	AllPasswords      bool `json:"allPasswords"`
	UnlockedPasswords bool `json:"unlockedPasswords"`
}

func (s *Sites) UnmarshalJSON(reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return err
	}
	return nil
}
