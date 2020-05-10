package grpcserver

import (
	"context"
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

	Create(ctx context.Context, message *comcn.Message) (*comcn.Message, error)
	Read(ctx context.Context, message *comcn.Message) (*comcn.Message, error)
	Update(ctx context.Context, message *comcn.Message) (*comcn.Message, error)
	Delete(ctx context.Context, message *comcn.Message) (*comcn.Message, error)

	Get(ctx context.Context, message *comcn.Message) (*comcn.Message, error)
	Set(ctx context.Context, message *comcn.Message) (*comcn.Message, error)
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
