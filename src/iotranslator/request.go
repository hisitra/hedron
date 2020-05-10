package iotranslator

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hisitra/hedron/src/configs"
	"time"
)

type Request struct {
	ID    string `json:"ID"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Mode  string `json:"mode"`

	ArrivedAt time.Time `json:"arrivedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewRequest(data []byte, mode string) (*Request, *Response) {
	newReq := &Request{}
	err := json.Unmarshal(data, newReq)
	if err != nil {
		return nil, BadRequestResponse("Failed to Unmarshal")
	}
	newReq.Mode = mode

	if resp := newReq.validate(); !resp.IsSuccessful() {
		return nil, resp
	}

	newReq.ID = uuid.New().String()
	newReq.ArrivedAt = time.Now()
	newReq.ExpiresAt = newReq.ArrivedAt.Add(
		time.Duration(configs.Server.RequestTimeout) * time.Second)

	return newReq, SuccessResponse("")
}

func UnmarshalRequest(data []byte) (*Request, *Response) {
	req := &Request{}
	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, BadRequestResponse("Failed to unmarshal")
	}

	if resp := req.validate(); !resp.IsSuccessful() {
		return nil, resp
	}

	return req, SuccessResponse("")
}

func (r *Request) validate() *Response {
	// TODO: Validate key and value using Regex.
	if r.Key == "" {
		return BadRequestResponse("Empty key not allowed in Request.")
	}
	if r.Mode != "C" && r.Mode != "R" && r.Mode != "U" && r.Mode != "D" {
		return BadRequestResponse("Invalid Request Type: " + r.Mode)
	}
	return SuccessResponse("")
}

func (r *Request) Marshal() ([]byte, *Response) {
	reqJSON, err := json.Marshal(r)
	if err != nil {
		return nil, InternalServerResponse("")
	}

	return reqJSON, SuccessResponse("")
}
