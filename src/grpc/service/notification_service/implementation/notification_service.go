package implementation

import (
	"context"
	"notification-service/domain/enum"
	pb "notification-service/grpc/service/notification_service"
	"notification-service/usecase"
)

type NotificationServiceImpl struct {
	pb.UnimplementedNotificationServer
	NotificationUsecase usecase.NotificationUsecase
}


func NewNotificationServiceImpl(notificationUsecase usecase.NotificationUsecase) *NotificationServiceImpl {
	return &NotificationServiceImpl{}
}

func (n *NotificationServiceImpl) SendNotification(ctx context.Context, in *pb.NotificationMessage) (*pb.EmptyMessage, error) {

	notificationType := mapProtoNotificationtypeToNotificationType(in.NotificationType)
	n.NotificationUsecase.SendNotification(in.Sender, in.Receiver, notificationType, "")

	
	return &pb.EmptyMessage{}, nil
}


func mapProtoNotificationtypeToNotificationType(pbType pb.NotificationType) enum.NotificationType {

	if pbType == pb.NotificationType_Like {
		return enum.Like
	} else if pbType == pb.NotificationType_Dislike {
		return enum.Dislike
	} else if pbType == pb.NotificationType_Comment {
		return enum.Comment
	} else if pbType == pb.NotificationType_Post {
		return enum.Post
	} else if pbType == pb.NotificationType_Follow {
		return enum.Follow
	}

	return enum.Story


}