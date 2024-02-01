package client

import (
	"fmt"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	fetcherpb "golang-project-template/internal/datafetcher/ports/proto/pb"

	managerpb "golang-project-template/internal/postmanager/ports/grpc/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientService struct {
	FetcherClient fetcherpb.SavePostsServiceClient
	ManagerClient managerpb.ManagePostsServiceClient
}

func NewClientService(fAddr, mAddr string) *ClientService {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	fetcherConn, err := grpc.Dial(fAddr, opts...)
	if err != nil {
		log.Printf("failed to dial on address %v\n, ERR: %v\n", fAddr, err)
	}
	fmt.Println("before fetcher client.....")
	FetcherClient := fetcherpb.NewSavePostsServiceClient(fetcherConn)

	managerConn, err := grpc.Dial(mAddr, opts...)
	if err != nil {
		log.Printf("failed to dial on address %v\n, ERR: %v\n", mAddr, err)
	}
	ManagerClient := managerpb.NewManagePostsServiceClient(managerConn)

	return &ClientService{
		FetcherClient: FetcherClient,
		ManagerClient: ManagerClient,
	}
}
