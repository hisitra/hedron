package main

import (
	"context"
	protocom "github.com/hisitra/hedron/src/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := protocom.NewExternalClient(conn)

	_, err = client.Get(context.Background(), &protocom.ExternalGetRequest{
		Data: nil,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.Set(context.Background(), &protocom.ExternalSetRequest{
		Data: nil,
	})
	if err != nil {
		panic(err)
	}
}
