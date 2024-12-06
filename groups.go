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

type NewGroupRequest struct {
	Name string `json:"name"`
}

type NewGroupResponse struct {
	ID string `json:"id"`
}

type EditGroupRequest struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

type DeleteGroupRequest struct {
	ID string `json:"id"`
}

func (a *AltaClient) AddGroup(name string) (*string, error) {
	URL := "group/add"

	var req = NewGroupRequest{Name: name}

	var resp NewGroupResponse

	if err := a.postRequest(URL, req, &resp); err != nil {
		return nil, fmt.Errorf("failed to add group: %w", err)
	}

	return &resp.ID, nil
}

func (a *AltaClient) EditGroup(req EditGroupRequest) error {
	URL := "group/edit"

	if err := a.postRequest(URL, req, nil); err != nil {
		return fmt.Errorf("failed to edit group: %w", err)
	}

	return nil
}

func (a *AltaClient) DeleteGroup(id string) error {
	URL := "group/delete"

	req := DeleteGroupRequest{
		ID: id,
	}

	if err := a.postRequest(URL, req, nil); err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}
