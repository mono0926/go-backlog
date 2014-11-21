package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	teamKey := "TODO"
	apiKey := "TODO"
	notifications := getNotifications(teamKey, apiKey)
	notifications = collectUnreadNotifications(notifications)
	if len(notifications) == 0 {
		return
	}
	fmt.Printf("unread count: %d", len(notifications))
	for _, n := range notifications {
		fmt.Printf("comment: %s (%s)", n.Comment.Content, n.Created)
	}
}

func parseDateString(d string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", d)
	return t
}

func collectUnreadNotifications(notifications []NotificationResponse) []NotificationResponse {
	r := make([]NotificationResponse, 0)
	for _, n := range notifications {
		if !n.AlreadyRead {
			fmt.Println(n)
			r = append(r, n)
		}
	}
	return r
}

func getNotifications(teamKey string, apiKey string) []NotificationResponse {
	url := fmt.Sprintf("https://%s.backlog.jp/api/v2/notifications?apiKey=%s", teamKey, apiKey)
	r, _ := http.Get(url)
	defer r.Body.Close()
	var notifications []NotificationResponse
	bytes, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(bytes, &notifications)
	if error != nil {
		log.Fatalln(error)
	}
	return notifications
}

type NotificationResponse struct {
	AlreadyRead bool `json:"alreadyRead"`
	Comment     struct {
		ChangeLog []struct {
		AttachmentInfo   interface{} `json:"attachmentInfo"`
		AttributeInfo    interface{} `json:"attributeInfo"`
		Field            string      `json:"field"`
		NewValue         string      `json:"newValue"`
		NotificationInfo interface{} `json:"notificationInfo"`
		OriginalValue    string      `json:"originalValue"`
	} `json:"changeLog"`
		Content     string `json:"content"`
		Created     string `json:"created"`
		CreatedUser struct {
			ID          float64     `json:"id"`
			Lang        interface{} `json:"lang"`
			MailAddress string      `json:"mailAddress"`
			Name        string      `json:"name"`
			RoleType    float64     `json:"roleType"`
			UserId      string      `json:"userId"`
		} `json:"createdUser"`
		ID            float64 `json:"id"`
		Notifications []struct {
		AlreadyRead         bool    `json:"alreadyRead"`
		ID                  float64 `json:"id"`
		Reason              float64 `json:"reason"`
		ResourceAlreadyRead bool    `json:"resourceAlreadyRead"`
		User                struct {
			ID          float64 `json:"id"`
			Lang        string  `json:"lang"`
			MailAddress string  `json:"mailAddress"`
			Name        string  `json:"name"`
			RoleType    float64 `json:"roleType"`
			UserId      string  `json:"userId"`
		} `json:"user"`
	} `json:"notifications"`
		Stars   []interface{} `json:"stars"`
		Updated string        `json:"updated"`
	} `json:"comment"`
	Created string  `json:"created"`
	ID      float64 `json:"id"`
	Issue   struct {
		ActualHours interface{} `json:"actualHours"`
		Assignee    struct {
			ID          float64     `json:"id"`
			Lang        interface{} `json:"lang"`
			MailAddress string      `json:"mailAddress"`
			Name        string      `json:"name"`
			RoleType    float64     `json:"roleType"`
			UserId      string      `json:"userId"`
		} `json:"assignee"`
		Attachments []struct {
		Created     string `json:"created"`
		CreatedUser struct {
			ID          float64     `json:"id"`
			Lang        interface{} `json:"lang"`
			MailAddress string      `json:"mailAddress"`
			Name        string      `json:"name"`
			RoleType    float64     `json:"roleType"`
			UserId      string      `json:"userId"`
		} `json:"createdUser"`
		ID   float64 `json:"id"`
		Name string  `json:"name"`
		Size float64 `json:"size"`
	} `json:"attachments"`
		Category []struct {
		DisplayOrder float64 `json:"displayOrder"`
		ID           float64 `json:"id"`
		Name         string  `json:"name"`
	} `json:"category"`
		Created     string `json:"created"`
		CreatedUser struct {
			ID          float64     `json:"id"`
			Lang        interface{} `json:"lang"`
			MailAddress string      `json:"mailAddress"`
			Name        string      `json:"name"`
			RoleType    float64     `json:"roleType"`
			UserId      string      `json:"userId"`
		} `json:"createdUser"`
		CustomFields   []interface{} `json:"customFields"`
		Description    string        `json:"description"`
		DueDate        interface{}   `json:"dueDate"`
		EstimatedHours interface{}   `json:"estimatedHours"`
		ID             float64       `json:"id"`
		IssueKey       string        `json:"issueKey"`
		IssueType      struct {
			Color        string  `json:"color"`
			DisplayOrder float64 `json:"displayOrder"`
			ID           float64 `json:"id"`
			Name         string  `json:"name"`
			ProjectId    float64 `json:"projectId"`
		} `json:"issueType"`
		KeyId     float64 `json:"keyId"`
		Milestone []struct {
		Archived       bool        `json:"archived"`
		Description    string      `json:"description"`
		DisplayOrder   float64     `json:"displayOrder"`
		ID             float64     `json:"id"`
		Name           string      `json:"name"`
		ProjectId      float64     `json:"projectId"`
		ReleaseDueDate interface{} `json:"releaseDueDate"`
		StartDate      interface{} `json:"startDate"`
	} `json:"milestone"`
		ParentIssueId float64 `json:"parentIssueId"`
		Priority      struct {
			ID   float64 `json:"id"`
			Name string  `json:"name"`
		} `json:"priority"`
		ProjectId  float64 `json:"projectId"`
		Resolution struct {
			ID   float64 `json:"id"`
			Name string  `json:"name"`
		} `json:"resolution"`
		SharedFiles []interface{} `json:"sharedFiles"`
		Stars       []interface{} `json:"stars"`
		StartDate   interface{}   `json:"startDate"`
		Status      struct {
			ID   float64 `json:"id"`
			Name string  `json:"name"`
		} `json:"status"`
		Summary     string `json:"summary"`
		Updated     string `json:"updated"`
		UpdatedUser struct {
			ID          float64     `json:"id"`
			Lang        interface{} `json:"lang"`
			MailAddress string      `json:"mailAddress"`
			Name        string      `json:"name"`
			RoleType    float64     `json:"roleType"`
			UserId      string      `json:"userId"`
		} `json:"updatedUser"`
		Versions []interface{} `json:"versions"`
	} `json:"issue"`
	Project struct {
		Archived           bool    `json:"archived"`
		ChartEnabled       bool    `json:"chartEnabled"`
		DisplayOrder       float64 `json:"displayOrder"`
		ID                 float64 `json:"id"`
		Name               string  `json:"name"`
		ProjectKey         string  `json:"projectKey"`
		SubtaskingEnabled  bool    `json:"subtaskingEnabled"`
		TextFormattingRule string  `json:"textFormattingRule"`
	} `json:"project"`
	Reason              float64 `json:"reason"`
	ResourceAlreadyRead bool    `json:"resourceAlreadyRead"`
	Sender              struct {
		ID          float64     `json:"id"`
		Lang        interface{} `json:"lang"`
		MailAddress string      `json:"mailAddress"`
		Name        string      `json:"name"`
		RoleType    float64     `json:"roleType"`
		UserId      string      `json:"userId"`
	} `json:"sender"`
}
