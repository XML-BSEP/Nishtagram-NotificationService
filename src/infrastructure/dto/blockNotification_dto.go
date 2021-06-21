package dto

type BlockNotificationDto struct {
	UserFor string `json:"user_for"`
	UserBy string `json:"user_by"`
	NotificationType NotificationTypeDto `json:"settings"`
}

type NotificationTypeDto struct {
	Type string `json:"type"`
	Value bool `json:"value"`
}
