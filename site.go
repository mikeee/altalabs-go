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

type Site struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Icon     string        `json:"icon"`
	Devices  []interface{} `json:"devices"` // TODO: implement devices struct
	Online   int           `json:"online"`
	Emails   []string      `json:"emails"`
	Identity struct {
		Email             string `json:"email"`
		Admin             bool   `json:"admin"`
		AllPasswords      bool   `json:"allPasswords"`
		UnlockedPasswords bool   `json:"unlockedPasswords"`
	} `json:"perms"`
}

type Sites map[string]Site

// TODO: unmarshal sites into a map by string rather than ID
