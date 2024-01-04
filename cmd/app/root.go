package app

import (
	"fmt"
	"golang-project-template/cmd/app/servers"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "run-grpc"}

// command to run grpc server
var userGrpcServerCmd = &cobra.Command{
	Use: "grpc-server",
	Run: func(cmd *cobra.Command, args []string) {
		servers.RunGrpcServer()
	},
}

func Execute() {
	rootCmd.AddCommand(userGrpcServerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
