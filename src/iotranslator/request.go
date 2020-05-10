package iotranslator

type Request struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func (r *Request) ToJSON() ([]byte, error) {
	return nil, nil
}
