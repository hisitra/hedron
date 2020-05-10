package executionmanager

import (
	iot "github.com/hisitra/hedron/src/iotranslator"
	"time"
)

type orderKeep struct {
	readsInProgress uint64
	writeInProgress bool

	waiting chan *iot.Request
	next    *iot.Request

	earliestReq      *iot.Request
	earliestWriteReq *iot.Request
}

func (o *orderKeep) isEmpty() bool {
	return !o.writeInProgress && o.readsInProgress == 0
}

func (o *orderKeep) updateTimes(req *iot.Request) {
	if !req.IsRead() &&
		(o.earliestWriteReq == nil || req.IsEarlierThan(o.earliestWriteReq)) {
		o.earliestWriteReq = req
	}
	if o.earliestReq == nil || req.IsEarlierThan(o.earliestReq) {
		o.earliestReq = req
	}
}

func (o *orderKeep) markCompletion(req *iot.Request) {
	if o.earliestReq.ID == req.ID {
		o.earliestReq = nil
	}
	if o.earliestWriteReq.ID == req.ID {
		o.earliestWriteReq = nil
	}
	if req.IsRead() && o.readsInProgress > 0 {
		o.readsInProgress -= 1
	}
	if !req.IsRead() {
		o.writeInProgress = false
	}
}

func (o *orderKeep) getNextExecutable() *iot.Request {
	req := o.getNextRequest()
	if req == nil {
		time.Sleep(time.Millisecond)
		req = o.getNextRequest()
		if req == nil {
			return nil
		}
	}

	if o.writeInProgress {
		return nil
	}

	if req.IsRead() {
		o.readsInProgress += 1
	} else if o.readsInProgress == 0 {
		o.writeInProgress = true
	} else {
		return nil
	}

	return o.next
}

func (o *orderKeep) getNextRequest() *iot.Request {
	if o.next != nil {
		return o.next
	}

	select {
	case nextReq, ok := <-o.waiting:
		if ok {
			o.next = nextReq
			return o.next
		}
	default:
		return nil
	}
	return nil
}
