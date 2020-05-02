package iotranslator

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"log"
	"time"
)

type Request struct {
	ID           string `json:"id"`
	Key          string `json:"key"`
	Value        string `json:"value,omitempty"`
	NodePassword string `json:"nodePassword"`

	Type      string    `json:"type"`
	ArrivedAt time.Time `json:"arrivedAt"`
	ExpiresAt time.Time `json:"expiresAt"`

	ExecutionStatus uint
	responseChan    chan *comcn.Output
}

func (r *Request) validate() *comcn.Output {
	if r.Key == "" {
		// TODO: Check key with a regex for invalid characters as well
		return BadRequestResponse("Empty key not allowed in request.")
	}
	if r.NodePassword != configs.Node.Password {
		return UnauthorizedResponse("Invalid NodePassword")
	}
	return SuccessResponse("")
}

func (r *Request) Encode() ([]byte, error) {
	reqJSON, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return reqJSON, nil
}

func NewRequest(inp *comcn.Input, mode string) (*Request, *comcn.Output) {
	req := &Request{
		Key:          inp.Key,
		Value:        inp.Value,
		NodePassword: inp.NodePassword,
	}

	res := req.validate()
	if res.Code != 200 {
		return nil, res
	}

	req.ID = uuid.New().String()
	req.Type = mode
	req.ExecutionStatus = 0
	req.responseChan = make(chan *comcn.Output)
	return req, SuccessResponse("")
}

func DecodeRequest(reqJSON []byte) (*Request, error) {
	req := &Request{}
	err := json.Unmarshal(reqJSON, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *Request) MarkAccepted() {
	r.ArrivedAt = time.Now()
	r.ExpiresAt = r.ArrivedAt.Add(
		time.Duration(configs.Server.RequestTimeout) * time.Second)
}

func (r *Request) IsRead() bool {
	return r.Type == "R"
}

func (r *Request) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}

func (r *Request) IsEarlierThan(anotherArrival time.Time) bool {
	return r.ArrivedAt.Before(anotherArrival)
}

func (r *Request) SendResponse(res *comcn.Output) {
	defer func(r *Request) {
		err := recover()
		if err != nil {
			log.Printf("Error occurred while sending response for Request: %s, error: %v\n", r.ID, err)
		}
	}(r)

	r.responseChan <- res
	r.ExecutionStatus = 2
}

func (r *Request) OnResponse() chan *comcn.Output {
	return r.responseChan
}
