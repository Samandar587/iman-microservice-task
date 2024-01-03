package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "run-grpc"}

// command to run grpc server
var userGrpcServerCmd = &cobra.Command{
	Use: "grpc-server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// command to run http server
var httpCmd = &cobra.Command{
	Use: "http-server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.AddCommand(userGrpcServerCmd)
	rootCmd.AddCommand(httpCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
