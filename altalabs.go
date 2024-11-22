package altalabs

import "net/http"

type Config struct {
	Token string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithToken(token string) *Config {
	c.Token = token
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
