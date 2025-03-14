package protocol

import (
	"net"
	"pie-fire-dire/generated/grpc/beefpb"

	"google.golang.org/grpc"
)

func ServeGRPC() error {
	var (
		config   = app.config
		handlers = app.handlers.grpc
	)

	server := grpc.NewServer()

	beefpb.RegisterBeefServiceServer(server, handlers.Beef)

	listener, err := net.Listen("tcp", ":"+config.Server.GRPCPort)
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
