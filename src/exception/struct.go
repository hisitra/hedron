package exception

import (
	"encoding/json"
)

type Exception struct {
	Code       int    `json:"code"`
	CustomCode string `json:"customCode"`
	Message    string `json:"message"`
}

func (ed Exception) Error() string {
	return ed.Message
}

func (ed Exception) ToJSON() ([]byte, error) {
	return json.Marshal(ed)
}

// Any other exception methods can be declared here
