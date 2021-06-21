package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"notification-service/infrastructure/dto"
	"notification-service/usecase"
)

type blockNotificationHandler struct {
	BlockNotificationUsecase usecase.BlockNotificationUsecase
}

type BlockNotificationHandler interface {
	Block(ctx *gin.Context)
}

func NewBlockNotificationHandler(blockNotificationUsecase usecase.BlockNotificationUsecase) BlockNotificationHandler {
	return &blockNotificationHandler{BlockNotificationUsecase: blockNotificationUsecase}
}

func (b *blockNotificationHandler) Block(ctx *gin.Context) {

	decoder := json.NewDecoder(ctx.Request.Body)

	var notificationDto dto.BlockNotificationDto

	if err := decoder.Decode(&notificationDto); err != nil {
		ctx.JSON(400, gin.H{"message" : "Error decoding body"})
		return
	}

	if err := b.BlockNotificationUsecase.Block(ctx, notificationDto.NotificationType, notificationDto.BlockedFor, notificationDto.BlockedBy); err != nil {
		ctx.JSON(400, gin.H{"message" : "Error blocking notifications"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Successfully blocked notification"})

}
