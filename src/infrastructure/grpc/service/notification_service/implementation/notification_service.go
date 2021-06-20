package implementation

import (
	"context"
	"io"
	"notification-service/domain/enum"
	"notification-service/infrastructure/grpc/service/follow_service"
	pb "notification-service/infrastructure/grpc/service/notification_service"
	"notification-service/usecase"
	pbFollow "notification-service/infrastructure/grpc/service/follow_service"
)

type NotificationServiceImpl struct {
	pb.UnimplementedNotificationServer
	NotificationUsecase usecase.NotificationUsecase
	FollowClient follow_service.FollowServiceClient
}


func NewNotificationServiceImpl(notificationUsecase usecase.NotificationUsecase, followClient follow_service.FollowServiceClient) *NotificationServiceImpl {
	return &NotificationServiceImpl{NotificationUsecase: notificationUsecase, FollowClient: followClient}
}

func (n *NotificationServiceImpl) SendNotification(ctx context.Context, in *pb.NotificationMessage) (*pb.EmptyMessage, error) {

	notificationType := mapProtoNotificationtypeToNotificationType(in.NotificationType)
	notification := n.NotificationUsecase.CreateNotification(in.Sender, in.Receiver, notificationType, in.RedirectPath)
	n.NotificationUsecase.SendNotification(ctx, notification)
	_ = n.NotificationUsecase.SaveNotification(ctx, notification)


	return &pb.EmptyMessage{}, nil
}

func (n *NotificationServiceImpl) SendNotifications(ctx context.Context, in *pb.MultipleNotificationsMessage) (*pb.EmptyMessage, error) {
	notificationType := mapProtoNotificationtypeToNotificationType(in.NotificationType)

	cli, _ := n.FollowClient.SendUsers(ctx, &pbFollow.User{UserId: in.SenderId})

	for {
		ret, err := cli.Recv()
		if err == io.EOF {
			break
		}
		notification := n.NotificationUsecase.CreateNotification(in.SenderId, ret.FollowerId, notificationType, in.RedirectPath)
		n.NotificationUsecase.SendNotification(ctx, notification)
		_ = n.NotificationUsecase.SaveNotification(ctx, notification)
	}

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