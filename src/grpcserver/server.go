package grpcserver

import (
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/logan"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	Port string
}

type Server interface {
	Start()

	// All RPC signatures here
}

func New() Server {
	return &server{Port: configs.Server.Port}
}

func (s *server) Start() {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	// register the servers here
	// example.RegisterExampleServer(grpcServer, s)

	logan.Info.Println("Starting gRPC Server at PORT:", s.Port)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
