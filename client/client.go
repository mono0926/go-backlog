package client

import (
	"fmt"
	"net/http"
	"log"
)

type Client struct {
	auth *Auth
}

func (c Client) Get(resource string) *http.Response {
	url := fmt.Sprintf("https://%s.backlog.jp/api/v2/%s?apiKey=%s", c.auth.TeamKey, resource, c.auth.ApiKey)
	log.Println(url)
	r, _ := http.Get(url)
	return r
}

func NewClient(auth *Auth) *Client {
	return &Client{auth: auth}
}
