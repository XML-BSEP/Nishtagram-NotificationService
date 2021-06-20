package seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"notification-service/domain"
	"notification-service/domain/enum"
	"time"
)

func DropDatabase(db string, mongoCli *mongo.Client, ctx *context.Context) error{
	err := mongoCli.Database(db).Drop(*ctx)
	if err != nil {
		return err
	}

	return nil
}


func SeedData(db string, mongoCli *mongo.Client, ctx *context.Context) error{
	if err := DropDatabase(db, mongoCli, ctx); err != nil {
		return err
	}

	if cnt, _ := mongoCli.Database(db).Collection("notifications").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		notificationCollection := mongoCli.Database(db).Collection("notifications")
		seedNotifications(notificationCollection, ctx)
	}

	return nil
}

func seedNotifications(notifications *mongo.Collection, ctx *context.Context) {
	notification1From := domain.Profile{Id: "e2b5f92e-c31b-11eb-8529-0242ac130003"}
	notification1To := domain.Profile{Id: "23ddb1dd-4303-428b-b506-ff313071d5d7"}

	_, err := notifications.InsertMany(*ctx, []interface{} {
		bson.D{
			{"_id", "93b58288-dc97-4386-b5c1-0c4d8bf07e41"},
			{"timestamp", time.Now()},
			{"content", "User liked your photo"},
			{"redirect_path", "/"},
			{"read", false},
			{"type", enum.NotificationType(1)},
			{"notification_from", notification1From},
			{"notification_to", notification1To},
			},
	})

	if err != nil {
		log.Fatal(err)
	}
}
