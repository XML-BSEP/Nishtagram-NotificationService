package mapper

import (
	"notification-service/domain/enum"
	"notification-service/infrastructure/dto"
)

func NotificationTypeDtoToNotificationType(dto dto.NotificationTypeDto) enum.NotificationType {
	if dto.Type == "like" {
		return enum.NotificationType(0)
	}else if dto.Type == "dislike" {
		return enum.NotificationType(1)
	}else if dto.Type == "comment" {
		return enum.NotificationType(2)
	}else if dto.Type == "post" {
		return enum.NotificationType(3)
	}else if dto.Type == "follow" {
		return enum.NotificationType(4)
	}else {
		return enum.NotificationType(5)
	}


}

func NotificationTypeToNotificationTypeDto(notifications []enum.NotificationType) []dto.NotificationTypeDto {
	var notificationTypes []dto.NotificationTypeDto
	//notificationTypes := make([]dto.NotificationTypeDto, len(notifications))

	for _, it := range notifications {
		if it == 0 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "like", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
		if it == 1 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "dislike", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
		if it == 2 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "comment", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
		if it == 3 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "post", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
		if it == 4 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "follow", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
		if it == 5 {
			notificationTypeDto := dto.NotificationTypeDto{Type: "story", Value: true}
			notificationTypes = append(notificationTypes, notificationTypeDto)
		}
	}

	return notificationTypes
}