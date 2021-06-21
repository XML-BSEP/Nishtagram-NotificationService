package domain

import (
	"github.com/google/uuid"
	"notification-service/domain/enum"
	"time"
)

type Notification struct {
	Id		string`bson:"_id" json:"id"`
	Timestamp		time.Time `bson:"timestamp" json:"timestamp"`
	Content		string `bson:"content" json:"content"`
	RedirectPath	string `bson:"redirect_path" json:"redirect_path"`
	Read 	bool `bson:"read" json:"read"`
	Type 	enum.NotificationType `bson:"type" json:"type"`
	NotificationFrom Profile `bson:"notification_from" json:"notification_from"`
	NotificationTo Profile `bson:"notification_to" json:"notification_to"`
	SenderUsername string `bson:"sender_username" json:"sender_username"`
}

func NewNotification(content string, redirectPath string, read bool, notificationType enum.NotificationType, notificationFrom string, notificationTo string) Notification {
	notification := Notification{
		Id: uuid.NewString(),
		Read: read,
		RedirectPath: redirectPath,
		Type: notificationType,
		NotificationTo: Profile{Id: notificationTo},
		NotificationFrom: Profile{Id: notificationFrom},
		Content: content,
		Timestamp: time.Now(),
	}

	return notification
}