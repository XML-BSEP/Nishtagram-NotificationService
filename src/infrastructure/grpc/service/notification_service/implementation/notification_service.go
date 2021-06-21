package implementation

import (
	"context"
	"fmt"
	"io"
	"notification-service/domain/enum"
	"notification-service/infrastructure/grpc/service/follow_service"
	pbFollow "notification-service/infrastructure/grpc/service/follow_service"
	pb "notification-service/infrastructure/grpc/service/notification_service"
	"notification-service/infrastructure/grpc/service/user_service"
	"notification-service/usecase"
	pbUser "notification-service/infrastructure/grpc/service/user_service"
)

type NotificationServiceImpl struct {
	pb.UnimplementedNotificationServer
	NotificationUsecase usecase.NotificationUsecase
	FollowClient follow_service.FollowServiceClient
	BlockNotificationUsecase usecase.BlockNotificationUsecase
	UserClient user_service.UserDetailsClient
}


func NewNotificationServiceImpl(notificationUsecase usecase.NotificationUsecase, followClient follow_service.FollowServiceClient, blockNotificationUsecase usecase.BlockNotificationUsecase, userClient user_service.UserDetailsClient) *NotificationServiceImpl {
	return &NotificationServiceImpl{NotificationUsecase: notificationUsecase, FollowClient: followClient, BlockNotificationUsecase: blockNotificationUsecase, UserClient: userClient}
}

func (n *NotificationServiceImpl) SendNotification(ctx context.Context, in *pb.NotificationMessage) (*pb.EmptyMessage, error) {

	if in.Sender == in.Receiver {
		return nil, fmt.Errorf("")
	}
	isBlocked, _ := n.BlockNotificationUsecase.IsBlocked(ctx, in.Sender, in.Receiver)

	if isBlocked {
		return nil, fmt.Errorf("")
	}


	senderUsername, _ := n.UserClient.GetUsername(ctx, &pbUser.UserId{UserId: in.Sender})
	notificationType := mapProtoNotificationtypeToNotificationType(in.NotificationType)
	notification := n.NotificationUsecase.CreateNotification(in.Sender, in.Receiver, notificationType, in.RedirectPath, senderUsername.Username)



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

		if err != nil {
			break
		}
		if in.SenderId == ret.FollowerId {
			break
		}
		isBlocked, _ := n.BlockNotificationUsecase.IsBlocked(ctx, in.SenderId, ret.FollowerId)

		if isBlocked {
			break
		}
		senderUsername, _ := n.UserClient.GetUsername(ctx, &pbUser.UserId{UserId: in.SenderId})
		notification := n.NotificationUsecase.CreateNotification(in.SenderId, ret.FollowerId, notificationType, in.RedirectPath, senderUsername.Username)
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