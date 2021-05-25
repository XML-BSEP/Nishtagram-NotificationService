package domain

import (
	"notification-service/domain/enum"
	"time"
)

type Notification struct {
	Id		uint64`json:"id"`
	Timestamp		time.Time `json:"timestamp"`
	Content		string `json:"content"`
	RedirectPath	string `json:"redirect_path"`
	Read 	bool `json:"read"`
	Type 	enum.NotificationType `json:"type"`
}