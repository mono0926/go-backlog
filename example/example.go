package main

import (
	"fmt"
	"github.com/mono0926/go-backlog/api"
)

func main() {
	teamKey := "YOUR_TEAM_KEY"
	apiKey := "YOUR_API_KEY"
	a := api.NewApi(teamKey, apiKey)
	nc := a.NotificationClient()
	notifications := nc.GetNotifications(true)
	if len(notifications) == 0 {
		return
	}
	fmt.Printf("unread count: %d", len(notifications))
	for _, n := range notifications {
		fmt.Printf("comment: %s (%s)", n.Comment.Content, n.Created)
	}
}
