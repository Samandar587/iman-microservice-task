package servers

import (
	"fmt"
	common "golang-project-template/internal/common/db"
	"golang-project-template/internal/datafetcher/adapters"
	"golang-project-template/internal/datafetcher/app"
	dataFetcherServer "golang-project-template/internal/datafetcher/ports"
	dataFetcherPb "golang-project-template/internal/datafetcher/ports/proto/pb"
	"golang-project-template/internal/pkg/config"
	"net/http"
	"os"

	apiGatewayApp "golang-project-template/internal/apigateway/app"
	"golang-project-template/internal/apigateway/pgk/client"
	apiGateWayPorts "golang-project-template/internal/apigateway/ports"
	grpcCommon "golang-project-template/internal/common/grpc"
	postManagerAdapers "golang-project-template/internal/postmanager/adapters"
	postManagerApp "golang-project-template/internal/postmanager/app"
	postManagerServer "golang-project-template/internal/postmanager/ports/grpc"
	postManagerPb "golang-project-template/internal/postmanager/ports/grpc/proto/pb"

	chi "github.com/go-chi/chi/v5"

	"log"

	"google.golang.org/grpc"
)

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

func ApiGatewayServer() {
	fetcherAddr := os.Getenv("DATAFETCHER_ADRESS")
	managerAddr := os.Getenv("POSTMANAGER_ADRESS")

	clients := client.NewClientService(fetcherAddr, managerAddr)
	app := apiGatewayApp.NewUsecase(clients.FetcherClient, clients.ManagerClient)
	controller := apiGateWayPorts.NewController(app)

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Get("/fetch", controller.CollectPostsHandler)
		r.Get("/getById/{id}", controller.GetByIdHandler)
		r.Post("/create", controller.CreateHandler)
		r.Put("/update", controller.UpdateHandler)
		r.Delete("/delete", controller.DeleteHandler)
	})

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting HTTP server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()
}
