package protocol

import (
	"net"
	"pie-fire-dire/generated/grpc/beefpb"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
}

func NewGRPCServer() *GRPCServer {
	var (
		config  = app.config
		handler = app.handlers.grpc
	)

	server := grpc.NewServer()
	beefpb.RegisterBeefServiceServer(server, handler.Beef)

	return &GRPCServer{
		server: server,
		port:   config.Server.GRPCPort,
	}
}

func (s *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *GRPCServer) Stop() error {
	s.server.GracefulStop()
	return nil
}
