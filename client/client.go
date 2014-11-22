package client

import "fmt"

type Client struct {
	auth          *Auth
}

func (c Client) ConstructURL(resource string) string {
	return fmt.Sprintf("https://%s.backlog.jp/api/v2/%s?apiKey=%s", c.auth.TeamKey, resource, c.auth.ApiKey)
}

func NewClient(auth *Auth) *Client {
	return &Client{auth: auth}
}
