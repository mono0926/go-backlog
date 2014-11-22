package client

type Client struct {
	Auth *Auth
}

func NewClient(auth *Auth) *Client {
	return &Client{Auth: auth}
}
