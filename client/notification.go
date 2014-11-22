package client

import (
	"encoding/json"
	"fmt"
	"github.com/mono0926/go-backlog/response"
	"io/ioutil"
	"log"
)

type Notification interface {
	GetNotifications(unreadOnly bool) []response.Notification
}

func NewNotification(auth *Auth) Notification {
	return &notification{Client: NewClient(auth)}
}

type notification struct {
	*Client
}

func (n notification) GetNotifications(unreadOnly bool) []response.Notification {
	r := n.Get("notifications")
	defer r.Body.Close()
	var notifications []response.Notification
	bytes, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(bytes, &notifications)
	if error != nil {
		log.Fatalln(error)
	}
	if !unreadOnly {
		return notifications
	}
	return filterUnread(notifications)
}

func filterUnread(notifications []response.Notification) []response.Notification {
	r := make([]response.Notification, 0)
	for _, n := range notifications {
		if !n.AlreadyRead {
			fmt.Println(n)
			r = append(r, n)
		}
	}
	return r
}
