package usecase

import (
	"context"
	"fmt"
	"notification-service/domain/enum"
	"notification-service/repository"
)

type blockNotificationUsecase struct {
	BlockNotificationRepository repository.BlockNotificationRepository
}


type BlockNotificationUsecase interface {
	IsBlocked(context context.Context, sender, receiver string) (bool, error)
	Block(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error
	GetBlockedTypes(context context.Context, blockedBy, blockedFor string) ([]enum.NotificationType, error)
	Unblock(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error
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

func (b *blockNotificationUsecase) GetBlockedTypes(context context.Context, blockedBy, blockedFor string) ([]enum.NotificationType, error) {
	return b.BlockNotificationRepository.GetBlockedTypes(context, blockedBy, blockedFor)
}

func (b *blockNotificationUsecase) Unblock(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error {
	stuff,err := b.BlockNotificationRepository.Unblock(context, notificationType, blockedBy, blockedFor)
	if err!=nil{
		return err
	}
	if stuff.DeletedCount==0{
		return fmt.Errorf("Error deleting one")
	}
	return nil
}

