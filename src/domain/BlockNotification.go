package domain

import (
	"notification-service/domain/enum"
	"time"
)

type BlockNotification struct {
	Id		uint64`json:"id"`
	Timestamp	time.Time `json:"timestamp"`
	Type	enum.BlockNotificationType	`json:"type"`
	BlockNotificationFor 	Profile `json:"block_notification_for"`
	BlockedNotification 	Profile `json:"blocked_notification"`

}
