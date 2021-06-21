package dto

import "notification-service/domain/enum"

type BlockNotificationDto struct {
	BlockedFor string `json:"blocked_for"`
	BlockedBy string `json:"blocked_by"`
	NotificationType enum.NotificationType `json:"notification_type"`
}
