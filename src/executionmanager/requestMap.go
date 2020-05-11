package executionmanager

import (
	"github.com/hisitra/hedron/src/configs"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"github.com/hisitra/ventager"
	"log"
	"sync"
)

var reqMap = &requestMap{
	data:   map[string]*orderKeep{},
	locker: sync.RWMutex{},
}

type requestMap struct {
	data   map[string]*orderKeep
	locker sync.RWMutex
}

func (rm *requestMap) acceptRequest(req *iot.Request) {
	rm.locker.Lock()
	defer rm.locker.Unlock()

	keep, exists := rm.data[req.Key]
	if !exists {
		keep = &orderKeep{
			key:     req.Key,
			waiting: make(chan *iot.Request),
		}
	}

	keep.addRequest(req)
	rm.drainKeep(keep)
}

func (rm *requestMap) removeRequest(req *iot.Request) {
	rm.locker.Lock()
	rm.locker.Unlock()

	keep, exists := rm.data[req.Key]
	if !exists {
		log.Println("Warning: Attempt to remove a request from non-existing orderKeep.")
		return
	}
	keep.removeRequest(req)
}

func (rm *requestMap) drainKeep(keep *orderKeep) {
	for {
		nextReq := keep.getNext()
		if nextReq == nil {
			log.Printf("Info: Waiting queue for Key: %s exhausted.", keep.key)
			return
		}

		if !keep.isExecutable(nextReq) {
			return
		}
		keep.setNext(nil)
		ventager.Fire(configs.Events.RequestExecutionBegin, nextReq)
	}
}

// TODO: Add IsExecutable method for internal GET requests.
