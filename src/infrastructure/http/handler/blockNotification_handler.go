package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"notification-service/infrastructure/dto"
	"notification-service/infrastructure/mapper"
	"notification-service/usecase"
)

type blockNotificationHandler struct {
	BlockNotificationUsecase usecase.BlockNotificationUsecase
}


type BlockNotificationHandler interface {
	Block(ctx *gin.Context)
	GetBlockedTypes(ctx *gin.Context)
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

	notificationType := mapper.NotificationTypeDtoToNotificationType(notificationDto.NotificationType)

	if err := b.BlockNotificationUsecase.Block(ctx, notificationType, notificationDto.UserBy, notificationDto.UserFor); err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not block notification"})
		return
	}


	ctx.JSON(200, gin.H{"message" : "Successfully blocked notification"})

}

func (b *blockNotificationHandler) GetBlockedTypes(ctx *gin.Context) {

	blockedBy := ctx.Param("blockedBy")
	blockedFor := ctx.Param("blockedFor")

	blockedTypes, err := b.BlockNotificationUsecase.GetBlockedTypes(ctx, blockedBy, blockedFor)

	if err != nil {
		ctx.JSON(204, gin.H{"message" : "There is no blocked notifications"})
		return
	}

	blockedTypesDto := mapper.NotificationTypeToNotificationTypeDto(blockedTypes)

	ctx.JSON(200, blockedTypesDto)
}
