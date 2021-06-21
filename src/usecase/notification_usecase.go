package usecase

import (
	"context"
	"github.com/google/uuid"
	"notification-service/domain"
	"notification-service/domain/enum"
	"notification-service/infrastructure/pusher"
	"notification-service/repository"
	"time"
)

type notificationUsecase struct {
	PusherService pusher.PusherService
	NotificationRepository repository.NotificationRepository
}


type NotificationUsecase interface {
	SendNotification(context context.Context, notification domain.Notification)
	GetNotificationsForUser(context context.Context, userId string) (*[]domain.Notification, error)
	UpdateNotificationStatus(context context.Context, notificationId string, status bool) error
	SaveNotification(context context.Context, notification domain.Notification) error
	CreateNotification(sender string, receiver string, notificationType enum.NotificationType, redirectPath string, senderUsername string) domain.Notification
}

func NewNotificationUsecase(pusherService pusher.PusherService, notificationRepository repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{PusherService: pusherService, NotificationRepository: notificationRepository}
}

func (n *notificationUsecase) SendNotification(context context.Context, notification domain.Notification) {

	go n.PusherService.Trigger(notification.NotificationTo.Id, "notification", notification)
}


func (n *notificationUsecase) CreateNotification(sender string, receiver string, notificationType enum.NotificationType, redirectPath string, senderUsername string) domain.Notification {


	id := uuid.NewString()
	timestamp := time.Now()
	notificationContent := n.CreateNotificationContent(senderUsername, notificationType)


	notification := domain.Notification{
		Id: id,
		Timestamp: timestamp,
		Content: notificationContent,
		RedirectPath: redirectPath,
		Read: false,
		Type: notificationType,
		NotificationFrom: domain.Profile{Id: sender},
		NotificationTo: domain.Profile{Id: receiver},
		SenderUsername: senderUsername,
	}

	return notification
}

func (n *notificationUsecase) GetNotificationsForUser(context context.Context, userId string) (*[]domain.Notification, error) {
	return n.NotificationRepository.GetNotificationsForUser(context, userId)
}

func (n *notificationUsecase) UpdateNotificationStatus(context context.Context, notificationId string, status bool) error {
	return n.NotificationRepository.UpdateNotificationStatus(context, notificationId, status)
}

func (n *notificationUsecase) SaveNotification(context context.Context, notification domain.Notification) error {
	return n.NotificationRepository.SaveNotification(context, notification)
}

func (n *notificationUsecase) CreateNotificationContent(sender string, notificationType enum.NotificationType) string {
	if notificationType == enum.Like {
		return sender + " liked your photo"
	}

	if notificationType == enum.Dislike {
		return sender + " disliked your photo"
	}

	if notificationType == enum.Comment {
		return sender + " commented your post"
	}

	if notificationType == enum.Post {
		return sender + " added new post"
	}

	return ""
}
