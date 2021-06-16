package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"notification-service/grpc/service/notification_service"
	pusher2 "notification-service/infrastructure/pusher"
	"notification-service/interactor"

)

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func main() {

	pusherClient := pusher2.GetConnection()
	if pusherClient == nil {
		panic("Can not create pusher client")
	}

	interactor := interactor.NewInteractor(pusherClient)

	port := uint(8078)
	list := getNetListener(port)

	creds, err := credentials.NewServerTLSFromFile("certificate/cert.pem", "certificate/key.pem")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	s := interactor.NewNotificationServiceImpl()
	notification_service.RegisterNotificationServer(grpcServer, s)

	fmt.Printf("Grpc server is listening on port: %d", port)
	log.Fatal(grpcServer.Serve(list))


}
