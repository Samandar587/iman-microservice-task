package app

import (
	"fmt"
	"golang-project-template/cmd/app/servers"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "run-grpc"}

// command to run grpc server
var dataFetcherGrpcCmd = &cobra.Command{
	Use: "grpc-server",
	Run: func(cmd *cobra.Command, args []string) {
		servers.RunDataFetcherGrpcServer()
	},
}

var postManagerGrpcCmd = &cobra.Command{
	Use: "post-manager-grpc-server",
	Run: func(cmd *cobra.Command, args []string) {
		servers.RunPostManagerGrpcServer()
	},
}

var apiGatewayCmd = &cobra.Command{
	Use: "api-gateway-server",
	Run: func(cmd *cobra.Command, args []string) {
		servers.ApiGatewayServer()
	},
}

func Execute() {
	rootCmd.AddCommand(dataFetcherGrpcCmd)
	rootCmd.AddCommand(postManagerGrpcCmd)
	rootCmd.AddCommand(apiGatewayCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
