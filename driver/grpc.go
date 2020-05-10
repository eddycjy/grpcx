package driver

import (
	"google.golang.org/grpc"
	"net"
)

type GRPCIface interface {
	DriverIface
	GetEngine() *grpc.Server
}

type gRPCServer struct {
	s *grpc.Server
	lis net.Listener
}

func NewGRPCServer(opt ...grpc.ServerOption) GRPCIface {
	s := grpc.NewServer(opt...)
	return &gRPCServer{s: s}
}

func (grpc *gRPCServer) GetEngine() *grpc.Server {
	return grpc.s
}

func (grpc *gRPCServer) SetListener(lis net.Listener) {
	grpc.lis = lis
}

func (grpc *gRPCServer) Serve() error {
	return grpc.s.Serve(grpc.lis)
}

