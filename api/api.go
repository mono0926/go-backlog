package api

import "github.com/mono0926/go-backlog/client"

type Api interface {
	NotificationClient() client.Notification
}

func NewApi(teamKey string, apiKey string) Api {
	auth := client.NewAuth(teamKey, apiKey)
	return &api{auth: auth}
}

type api struct {
	auth *client.Auth
}

func (api api) NotificationClient() client.Notification {
	return client.NewNotification(api.auth)
}
