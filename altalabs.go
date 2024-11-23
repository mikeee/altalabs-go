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
