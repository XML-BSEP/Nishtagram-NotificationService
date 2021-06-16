package domain

import (
	"notification-service/domain/enum"
	"time"
)

type Notification struct {
	Id		string`json:"id"`
	Timestamp		time.Time `json:"timestamp"`
	Content		string `json:"content"`
	RedirectPath	string `json:"redirect_path"`
	Read 	bool `json:"read"`
	Type 	enum.NotificationType `json:"type"`
	NotificationFrom Profile `json:"notification_from"`
	NotificationTo Profile `json:"notification_to"`
}