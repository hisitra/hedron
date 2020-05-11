package executionmanager

import (
	iot "github.com/hisitra/hedron/src/iotranslator"
)

type orderKeep struct {
	key string

	readsInProgress uint64
	writeInProgress bool

	waiting chan *iot.Request
	next    *iot.Request

	earliestReq      *iot.Request
	earliestWriteReq *iot.Request
}
