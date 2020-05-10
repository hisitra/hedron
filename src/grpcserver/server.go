package grpcserver

import (
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	Port string
}

type Server interface {
	Start()
}

func New() Server {
	return &server{Port: configs.Server.RpcPort}
}

func (s *server) Start() {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	comcn.RegisterExternalServer(grpcServer, s)
	comcn.RegisterInternalServer(grpcServer, s)

	log.Println("Starting Hedron Node:", configs.Node.Name, "gRPC Server at PORT:", configs.Server.RpcPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
