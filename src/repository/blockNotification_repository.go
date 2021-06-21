package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"notification-service/domain"
	"notification-service/domain/enum"
	"time"
)

type blockNotificationRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}


type BlockNotificationRepository interface {
	IsBlocked(context context.Context, sender, receiver string) (bool, error)
	Block(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error
}

func NewBlockNotificationRepository(db *mongo.Client) BlockNotificationRepository {
	return &blockNotificationRepository{
		db : db,
		collection : db.Database("notification_db").Collection("block_notifications"),
	}
}

func (b *blockNotificationRepository) IsBlocked(context context.Context, sender, receiver string) (bool, error) {

	var blocked domain.BlockNotification
	err := b.collection.FindOne(context, bson.M{"blocked_by._id" : receiver, "blocked_for._id" : sender}).Decode(&blocked)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *blockNotificationRepository) Block(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) error {

	blockNotification := domain.BlockNotification{
		ID: uuid.NewString(),
		NotificationType: notificationType,
		Timestamp: time.Now(),
		BlockedBy: domain.Profile{Id: blockedBy},
		BlockedFor: domain.Profile{Id: blockedFor},
	}

	_, err := b.collection.InsertOne(context, blockNotification)

	if err != nil {
		return err
	}

	return nil
}




