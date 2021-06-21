package usecase

import (
	"context"
	"notification-service/domain/enum"
	"notification-service/repository"
)

type blockNotificationUsecase struct {
	BlockNotificationRepository repository.BlockNotificationRepository
}


type BlockNotificationUsecase interface {
	IsBlocked(context context.Context, sender, receiver string) (bool, error)
	Block(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error
}

func NewBlockNotificationUsecase(blockNotificationRepository repository.BlockNotificationRepository) BlockNotificationUsecase {
	return &blockNotificationUsecase{BlockNotificationRepository: blockNotificationRepository}
}

func (b *blockNotificationUsecase) IsBlocked(context context.Context, sender, receiver string) (bool, error) {
	return b.BlockNotificationRepository.IsBlocked(context, sender, receiver)
}

func (b *blockNotificationUsecase) Block(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error {
	return b.BlockNotificationRepository.Block(context, notificationType, blockedBy, blockedFor)
}