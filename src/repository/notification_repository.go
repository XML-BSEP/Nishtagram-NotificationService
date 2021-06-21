package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"notification-service/domain"
)

type notificationRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}

type NotificationRepository interface {
	GetNotificationsForUser(context context.Context, userId string) (*[]domain.Notification, error)
	UpdateNotificationStatus(context context.Context, notificationId string, status bool) error
	SaveNotification(context context.Context, notification domain.Notification) error
}

func NewNotificationRepository(db *mongo.Client) NotificationRepository {
	return &notificationRepository{
		db : db,
		collection: db.Database("notification_db").Collection("notifications"),
	}
}

func (n *notificationRepository) GetNotificationsForUser(context context.Context, userId string) (*[]domain.Notification, error) {

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	filter, err := n.collection.Find(context, bson.M{"notification_to._id" : userId}, findOptions)

	if err != nil {
		return nil, err
	}

	var notifications []domain.Notification
	if err := filter.All(context, &notifications); err != nil {
		return nil, err
	}

	return &notifications, nil
}

func (n *notificationRepository) UpdateNotificationStatus(context context.Context, notificationId string, status bool) error {

	notification := bson.M{"_id" : notificationId}
	updateNotification := bson.M{"$set" : bson.M{
		"read" : status,
	}}

	_, err := n.collection.UpdateOne(context, notification, updateNotification)

	if err != nil {
		return err
	}

	return nil
}


func (n *notificationRepository) SaveNotification(context context.Context, notification domain.Notification) error {

	_, err := n.collection.InsertOne(context, notification)

	if err != nil {
		return err
	}
	return nil
}

