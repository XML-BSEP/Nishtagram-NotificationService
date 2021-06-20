package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"notification-service/infrastructure/dto"
	pb "notification-service/infrastructure/grpc/service/follow_service"
	client2 "notification-service/infrastructure/grpc/service/follow_service/client"
	"notification-service/usecase"
)

type NotificationHandler interface {
	GetNotificationsByUserId(ctx *gin.Context)
	UpdateNotificationStatus(ctx *gin.Context)
	GrpcProba(ctx *gin.Context)
}

type notificationHandler struct {
	NotificationUsecase usecase.NotificationUsecase
}


func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) NotificationHandler {
	return &notificationHandler{NotificationUsecase: notificationUsecase}
}

func (n *notificationHandler) GetNotificationsByUserId(ctx *gin.Context) {

	userId := ctx.Param("userId")

	notifications, err := n.NotificationUsecase.GetNotificationsForUser(ctx, userId)

	if err != nil {
		ctx.JSON(404, gin.H{"message" : "There are no notifications for user"})
		return
	}

	ctx.JSON(200, notifications)

}

func (n *notificationHandler) UpdateNotificationStatus(ctx *gin.Context) {

	notificationId := ctx.Param("notificationId")

	var updateNotificationDto dto.UpdateNotificationDto

	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&updateNotificationDto); err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not decode body"})
		return
	}

	if err := n.NotificationUsecase.UpdateNotificationStatus(ctx, notificationId, updateNotificationDto.Read); err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not update notification status"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Successful updated"})

}

func (n *notificationHandler) GrpcProba(ctx *gin.Context) {

	client, err := client2.NewFollowClient("127.0.0.1:8077")

	if err != nil {
		ctx.JSON(400, gin.H{"message" : "Error client"})
		return
	}

	user := &pb.User{UserId: "e2b5f92e-c31b-11eb-8529-0242ac130003"}
	cli, err := client.SendUsers(ctx, user)

	var users []string
	//userIds := make([]string, 1000)

	for {
		in, err := cli.Recv()
		if err == io.EOF {
			break
		}

		ctx.JSON(200, in.FollowerId)

	}


	ctx.JSON(200, users)


}

