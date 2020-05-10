package iotranslator

import "time"

type Request struct {
	ID    string `json:"ID"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Mode  string `json:"mode"`

	ArrivedAt time.Time `json:"arrivedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (r *Request) ToJSON() ([]byte, error) {
	return nil, nil
}
