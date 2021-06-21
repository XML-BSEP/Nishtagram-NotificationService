package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"notification-service/infrastructure/grpc/interceptor"
	"notification-service/infrastructure/grpc/service/follow_service/client"
	"notification-service/infrastructure/grpc/service/notification_service"
	"notification-service/infrastructure/http/router"
	"notification-service/infrastructure/mongo"
	pusher2 "notification-service/infrastructure/pusher"
	"notification-service/infrastructure/seeder"
	"notification-service/interactor"
	"os"
	"strconv"
)

func getNetListener(port uint) net.Listener {
	var domain string
	if os.Getenv("DOCKER_ENV") == "" {
		domain = "127.0.0.1"
	} else {
		domain = "notificationms"
	}
	domain = domain + ":" + strconv.Itoa(int(port))
	lis, err := net.Listen("tcp", domain)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func main() {

	mongoCli, ctx := mongo.NewMongoClient()

	db := mongo.GetDbName()

	if err := seeder.SeedData(db, mongoCli, ctx); err != nil {
		log.Fatal(err)
	}

	var notificationDomain string
	if os.Getenv("DOCKER_ENV") == "" {
		notificationDomain = "127.0.0.1:8077"
	} else {
		notificationDomain = "followms:8077"
	}
	followClient, err := client.NewFollowClient(notificationDomain)

	if err != nil {
		panic(err)
	}

	pusherClient := pusher2.GetConnection()
	if pusherClient == nil {
		panic("Can not create pusher client")
	}

	i := interactor.NewInteractor(pusherClient, mongoCli, followClient)
	appHandler := i.NewAppHandler()

	g := router.NewRoute(appHandler)
	g.Use(gin.Logger())
	g.Use(gin.Recovery())


	port := uint(8078)
	list := getNetListener(port)


	_ = interceptor.NewAuthUnaryInterceptor() //call authorization interceptor
	grpcServer := grpc.NewServer()
	s := i.NewNotificationServiceImpl()
	notification_service.RegisterNotificationServer(grpcServer, s)

	fmt.Printf("Grpc server is listening on port: %d", port)
	go func() {
		log.Fatalln(grpcServer.Serve(list))
	}()

	g.Run(":8094")


}
