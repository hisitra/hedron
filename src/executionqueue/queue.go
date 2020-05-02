package executionqueue

import (
	iot "github.com/hisitra/hedron/src/iotranslator"
	"sync"
)

type exQueue struct {
	locker sync.RWMutex
	groups []*iot.ReqGroup
}

var queue = &exQueue{
	locker: sync.RWMutex{},
	groups: []*iot.ReqGroup{},
}

func (eq *exQueue) addRequest(req *iot.Request) {
	var lastAddableGroup *iot.ReqGroup
	for i := len(eq.groups) - 1; i >= 0; i-- {
		if !eq.groups[i].IsReqAddable(req) {
			break
		}
		lastAddableGroup = eq.groups[i]
	}

	if lastAddableGroup == nil {
		lastAddableGroup = iot.NewReqGroup()
		eq.groups = append(eq.groups, lastAddableGroup)
	}

	req.MarkAccepted()
	// Error should not arise in this step as all
	// scenarios are covered above.
	_ = lastAddableGroup.AddReq(req)
}

func ExecutionAllowable(req *iot.Request) bool {
	queue.locker.RLock()
	defer queue.locker.RUnlock()

	for _, group := range queue.groups {
		reqArr, exists := group.Reqs[req.Key]
		if !exists {
			return true
		}
		for _, ownReq := range reqArr {
			if !ownReq.IsEarlierThan(req.ArrivedAt) {
				continue
			}
			return req.IsRead() && ownReq.IsRead()
		}
	}
	return true
}
