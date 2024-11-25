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

import "net/http"

type Method int

const (
	METHOD_TOKEN Method = iota
	METHOD_SRP
)

type Config struct {
	Method
	Token    string
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithToken(token string) *Config {
	c.Method = METHOD_TOKEN
	c.Token = token
	return c
}

func (c *Config) WithSRPAuth(username, password string) *Config {
	c.Method = METHOD_SRP
	c.Username = username
	c.Password = password
	return c
}

type altaClient struct {
	*http.Client
	*Config
}

func (a *altaClient) ListSites() (*Sites, error) {
	//TODO implement me
	panic("implement me")
}

type AltaClient interface {
	ListSites() (*Sites, error)
}

var _ AltaClient = &altaClient{}
