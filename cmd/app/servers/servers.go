package servers

import (
	"fmt"
	"golang-project-template/internal/common"
	"golang-project-template/internal/datafetcher/adapters"
	"golang-project-template/internal/datafetcher/app"
	server "golang-project-template/internal/datafetcher/ports/grpc"
	"golang-project-template/internal/datafetcher/ports/grpc/proto/pb"

	postManagerAdapers "golang-project-template/internal/postmanager/adapters"
	postManagerApp "golang-project-template/internal/postmanager/app"
	postManagerGrpc "golang-project-template/internal/postmanager/ports/grpc"
	postManagerPb "golang-project-template/internal/postmanager/ports/grpc/proto/pb"
	"net"

	"log"
	"os"

	"google.golang.org/grpc"
)

/*
// GRPC servers
func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {

}
*/

func RunDataFetcherGrpcServer() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := adapters.NewPostRepository(db)
	provider := adapters.NewPostProvider()
	usecase := app.NewPostUsecase(repo, provider)

	dataFetcherGrpcServer := server.NewDataFetcherServer(usecase)

	lisener, err := net.Listen("tcp", ":5006")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("listening on port :5006")

	s := grpc.NewServer()

	pb.RegisterSavePostsServiceServer(s, dataFetcherGrpcServer)
	if err := s.Serve(lisener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RunPostManagerGrpcServer() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := postManagerAdapers.NewpostRepsitory(db)
	usecase := postManagerApp.NewPostUsecase(repo)
	grpcHandler := postManagerGrpc.NewPostManagerGrpcServer(usecase)

	listener, err := net.Listen("tcp", ":5007")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("listening on port: 5007")
	s := grpc.NewServer()
	postManagerPb.RegisterManagePostsServiceServer(s, grpcHandler)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
