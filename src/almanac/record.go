package almanac

import (
	"encoding/json"
	"errors"
)

type Record struct {
	Key          string `json:"key"`
	Committed    string `json:"committed"`
	Uncommitted  string `json:"uncommitted"`
	DeleteStatus uint   `json:"deleteStatus"`
	Version      uint64 `json:"version"`
}

func (r *Record) validate() error {
	if r.Key == "" {
		return errors.New("empty key not allowed in a record")
	}
	return nil
}

func NewRecord(key string, value string) (*Record, error) {
	newRec := &Record{
		Key:          key,
		Committed:    "",
		Uncommitted:  value,
		DeleteStatus: 0,
		Version:      0,
	}
	err := newRec.validate()
	if err != nil {
		return nil, err
	}
	return newRec, nil
}

func EncodeRecord(rec *Record) ([]byte, error) {
	recJSON, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}
	return recJSON, nil
}

func DecodeRecord(data []byte) (*Record, error) {
	rec := &Record{}
	err := json.Unmarshal(data, rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}
