package interactor

import (
	"github.com/pusher/pusher-http-go/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"notification-service/infrastructure/grpc/service/follow_service"
	"notification-service/infrastructure/grpc/service/notification_service/implementation"
	"notification-service/infrastructure/http/handler"
	pusher2 "notification-service/infrastructure/pusher"
	"notification-service/repository"
	"notification-service/usecase"
)

type interactor struct {
	PusherClient *pusher.Client
	db *mongo.Client
	followClient follow_service.FollowServiceClient
}

type Interactor interface {
	NewPusherService() pusher2.PusherService

	NewNotificationUsecase() usecase.NotificationUsecase
	NewBlockNotificationUsecase() usecase.BlockNotificationUsecase

	NewNotificationRepository() repository.NotificationRepository
	NewBlockNotificationRepository() repository.BlockNotificationRepository

	NewNotificationServiceImpl() *implementation.NotificationServiceImpl

	NewAppHandler() AppHandler



}

func NewInteractor(pusherClient *pusher.Client, db *mongo.Client, followClient follow_service.FollowServiceClient) Interactor {
	return &interactor{PusherClient: pusherClient, db : db, followClient: followClient}
}

func (i *interactor) NewAppHandler() AppHandler {
	appHandler := &appHandler{}
	appHandler.NotificationHandler = handler.NewNotificationHandler(i.NewNotificationUsecase())
	return appHandler
}

func (i *interactor) NewPusherService() pusher2.PusherService {
	return pusher2.NewPusherService(i.PusherClient)
}

func (i *interactor) NewNotificationUsecase() usecase.NotificationUsecase {
	return usecase.NewNotificationUsecase(i.NewPusherService(), i.NewNotificationRepository())
}

func (i *interactor) NewNotificationServiceImpl() *implementation.NotificationServiceImpl {
	return implementation.NewNotificationServiceImpl(i.NewNotificationUsecase(), i.followClient)
}

func (i *interactor) NewNotificationRepository() repository.NotificationRepository {
	return repository.NewNotificationRepository(i.db)
}

type appHandler struct {
	handler.NotificationHandler
}

type AppHandler interface {
	handler.NotificationHandler
}

func (i *interactor) NewBlockNotificationUsecase() usecase.BlockNotificationUsecase {
	return usecase.NewBlockNotificationUsecase(i.NewBlockNotificationRepository())
}

func (i *interactor) NewBlockNotificationRepository() repository.BlockNotificationRepository {
	return repository.NewBlockNotificationRepository(i.db)
}


