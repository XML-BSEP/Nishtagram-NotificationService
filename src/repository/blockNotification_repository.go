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
	GetBlockedTypes(context context.Context, blockedBy, blockedFor string) ([]enum.NotificationType, error)
	Unblock(context context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) (*mongo.DeleteResult,error)
}

func NewBlockNotificationRepository(db *mongo.Client) BlockNotificationRepository {
	return &blockNotificationRepository{
		db : db,
		collection : db.Database("notification_db").Collection("block_notifications"),
	}
}

func (b *blockNotificationRepository) IsBlocked(context context.Context, sender, receiver string) (bool, error) {
	senderBson := bson.M{"_id" : sender}
	receiverBson := bson.M{"_id" : receiver}
	var blocked domain.BlockNotification
	err := b.collection.FindOne(context, bson.M{"blocked_by" : receiverBson, "blocked_for" : senderBson}).Decode(&blocked)

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

func (b *blockNotificationRepository) GetBlockedTypes(context context.Context, blockedBy, blockedFor string) ([]enum.NotificationType, error) {
	blockedByBson := bson.M{"_id" : blockedBy}
	blockedForBson := bson.M{"_id" : blockedFor}

	filter, err := b.collection.Find(context, bson.M{"blocked_by" : blockedByBson, "blocked_for" : blockedForBson})

	if err != nil {
		return nil, err
	}

	var blockedNotifications []domain.BlockNotification
	if err := filter.All(context, &blockedNotifications); err != nil {
		return nil, err
	}
	var notificationTypes []enum.NotificationType

	//notificationTypes := make([]enum.NotificationType, len(blockedNotifications))

	for _, it:= range blockedNotifications {
		notificationTypes = append(notificationTypes, it.NotificationType)
	}

	return notificationTypes, nil
}

func (b *blockNotificationRepository) Unblock(ctx context.Context, notificationType enum.NotificationType, blockedBy, blockedFor string) (*mongo.DeleteResult, error) {

	blockedByBson := bson.M{"_id" : blockedBy}
	blockedForBson := bson.M{"_id" : blockedFor}
	var blocked domain.BlockNotification
	_ = b.collection.FindOne(ctx, bson.M{"blocked_by" : blockedByBson, "blocked_for" : blockedForBson, "notification_type" : notificationType}).Decode(&blocked)

	//var notification domain.BlockNotification
	//
	//bsonBytes, _ := bson.Marshal(result)
	//err := bson.Unmarshal(bsonBytes, &notification)
	//if err != nil {
	//	return nil, err
	//}

	result1, _ := b.collection.DeleteOne(ctx, bson.M{"_id" :blocked.ID})


	return result1, nil
}




