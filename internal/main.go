package main

import (
	"context"
	"github.com/hisitra/hedron/src/logan"
	protocom "github.com/hisitra/hedron/src/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			logan.Warn.Println("Error while closing gRPC connection: ", err)
		}
	}()

	client := protocom.NewExternalClient(conn)

	_, err = client.Get(context.Background(), &protocom.ExternalGetRequest{})
	if err != nil {
		panic(err)
	}

	_, err = client.Set(context.Background(), &protocom.ExternalSetRequest{})
	if err != nil {
		panic(err)
	}
}
