package internalclient

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
	iot "github.com/hisitra/hedron/src/iotranslator"
)

func Get(address string, req *iot.Request) *comcn.Output {
	reqJSON, err := req.Encode()
	if err != nil {
		return iot.BadRequestResponse("")
	}

	client := connMap.getConn(address)
	if client == nil {
		return iot.InternalServerErrorResponse("Failed to create connection with: " + address)
	}

	res, err := client.Get(context.Background(), &comcn.InternalMessage{Value: reqJSON})
	if err != nil {
		connMap.refreshConn(address)
	}

	res, err = client.Get(context.Background(), &comcn.InternalMessage{Value: reqJSON})
	if err != nil {
		return iot.InternalServerErrorResponse("Failed to communicate with: " + address)
	}
	return res
}

func Set(address string, recordJSON []byte) *comcn.Output {
	client := connMap.getConn(address)
	if client == nil {
		return iot.InternalServerErrorResponse("Failed to create connection with: " + address)
	}

	res, err := client.Set(context.Background(), &comcn.InternalMessage{Value: recordJSON})
	if err != nil {
		connMap.refreshConn(address)
	}

	res, err = client.Set(context.Background(), &comcn.InternalMessage{Value: recordJSON})
	if err != nil {
		return iot.InternalServerErrorResponse("Failed to communicate with: " + address)
	}
	return res
}
