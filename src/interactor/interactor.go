package interactor

import (
	"github.com/pusher/pusher-http-go/v5"
	"notification-service/grpc/service/notification_service/implementation"
	pusher2 "notification-service/infrastructure/pusher"
	"notification-service/usecase"
)

type interactor struct {
	PusherClient *pusher.Client
}


type Interactor interface {
	NewPusherService() pusher2.PusherService

	NewNotificationUsecase() usecase.NotificationUsecase

	NewNotificationServiceImpl() *implementation.NotificationServiceImpl
}

func NewInteractor(pusherClient *pusher.Client) Interactor {
	return &interactor{PusherClient: pusherClient}
}

func (i *interactor) NewPusherService() pusher2.PusherService {
	return pusher2.NewPusherService(i.PusherClient)
}

func (i *interactor) NewNotificationUsecase() usecase.NotificationUsecase {
	return usecase.NewNotificationUsecase(i.NewPusherService())
}

func (i *interactor) NewNotificationServiceImpl() *implementation.NotificationServiceImpl {
	return implementation.NewNotificationServiceImpl(i.NewNotificationUsecase())
}

