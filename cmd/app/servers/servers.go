package servers

import (
	"google.golang.org/grpc"
)

// HTTP server
func RunHttpServer() {

}

// GRPC servers
func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {

}
