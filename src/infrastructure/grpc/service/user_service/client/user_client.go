package client

import (
	"google.golang.org/grpc"
	pb "notification-service/infrastructure/grpc/service/user_service"
)

func NewUserClient(address string) (pb.UserDetailsClient, error) {
	/*creds, err := credentials.NewClientTLSFromFile("certificate/cert.pem", "")
	if err != nil {
		return nil, err
	}*/
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	client := pb.NewUserDetailsClient(conn)
	return client, nil
}
