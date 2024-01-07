package servers

import (
	"fmt"
	common "golang-project-template/internal/common/db"
	"golang-project-template/internal/datafetcher/adapters"
	"golang-project-template/internal/datafetcher/app"
	dataFetcherServer "golang-project-template/internal/datafetcher/ports"
	dataFetcherPb "golang-project-template/internal/datafetcher/ports/proto/pb"
	"golang-project-template/internal/pkg/config"

	grpcCommon "golang-project-template/internal/common/grpc"
	postManagerAdapers "golang-project-template/internal/postmanager/adapters"
	postManagerApp "golang-project-template/internal/postmanager/app"
	postManagerServer "golang-project-template/internal/postmanager/ports/grpc"
	postManagerPb "golang-project-template/internal/postmanager/ports/grpc/proto/pb"

	"log"

	"google.golang.org/grpc"
)

/*
// GRPC servers
func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {

}
*/

func RunDataFetcherGrpcServer() {

	var dbInfo = config.NewDB()

	db, err := common.ConnectToDb(
		dbInfo.DB.Host,
		dbInfo.DB.Port,
		dbInfo.DB.Database,
		dbInfo.DB.User,
		dbInfo.DB.Password,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := adapters.NewPostRepository(db)
	provider := adapters.NewPostProvider()
	usecase := app.NewPostUsecase(repo, provider)

	addr := fmt.Sprintf(":%s", dbInfo.DataFetcherRPCPort)
	grpcCommon.RunGRPCServerOnAddr(addr, func(server *grpc.Server) {
		svc := dataFetcherServer.NewDataFetcherServer(usecase)
		dataFetcherPb.RegisterSavePostsServiceServer(server, svc)

	})
}

func RunPostManagerGrpcServer() {
	var dbInfo = config.NewDB()

	db, err := common.ConnectToDb(
		dbInfo.DB.Host,
		dbInfo.DB.Port,
		dbInfo.DB.Database,
		dbInfo.DB.User,
		dbInfo.DB.Password,
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := postManagerAdapers.NewpostRepsitory(db)
	usecase := postManagerApp.NewPostUsecase(repo)

	addr := fmt.Sprintf(":%s", dbInfo.PostManagerRPCPort)
	grpcCommon.RunGRPCServerOnAddr(addr, func(server *grpc.Server) {
		svc := postManagerServer.NewPostManagerGrpcServer(usecase)
		postManagerPb.RegisterManagePostsServiceServer(server, svc)
	})

}
