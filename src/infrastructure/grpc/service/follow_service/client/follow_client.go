package client

import (
	"google.golang.org/grpc"
	pb "notification-service/infrastructure/grpc/service/follow_service"
)

func NewFollowClient(address string) (pb.FollowServiceClient, error) {
	/*creds, err := credentials.NewClientTLSFromFile("certificate/cert.pem", "")
	if err != nil {
		return nil, err
	}*/
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	client := pb.NewFollowServiceClient(conn)
	return client, nil
}
