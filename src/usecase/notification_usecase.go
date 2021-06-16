package usecase

import (
	"github.com/google/uuid"
	"notification-service/domain"
	"notification-service/domain/enum"
	"notification-service/infrastructure/pusher"
	"time"
)

type notificationUsecase struct {
	PusherService pusher.PusherService
}

type NotificationUsecase interface {
	SendNotification(sender string, reciever string, notificationType enum.NotificationType, redirectPath string)
}

func NewNotificationUsecase(pusherService pusher.PusherService) NotificationUsecase {
	return &notificationUsecase{PusherService: pusherService}
}

func (n *notificationUsecase) SendNotification(sender string, receiver string, notificationType enum.NotificationType, redirectPath string) {
	notificationStruct := n.CreateNotification(sender, receiver, notificationType, "dsadsa")

	go n.PusherService.Trigger(receiver, "notification", notificationStruct)


}


func (n *notificationUsecase) CreateNotification(sender string, receiver string, notificationType enum.NotificationType, redirectPath string) domain.Notification {


	id := uuid.NewString()
	timestamp := time.Now()
	notificationContent := createNotificationContent(sender, notificationType)


	notification := domain.Notification{
		Id: id,
		Timestamp: timestamp,
		Content: notificationContent,
		RedirectPath: redirectPath,
		Read: false,
		Type: notificationType,
		NotificationFrom: domain.Profile{Id: sender},
		NotificationTo: domain.Profile{Id: receiver},
	}

	return notification
}


func createNotificationContent(sender string, notificationType enum.NotificationType) string {
	if notificationType == enum.Like {
		return sender + " liked your photo"
	}

	if notificationType == enum.Dislike {
		return sender + " disliked your photo"
	}

	if notificationType == enum.Comment {
		return sender + " commented your post"
	}

	return ""
}
