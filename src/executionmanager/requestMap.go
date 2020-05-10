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

	ordKeep, exists := rm.data[req.Key]
	if !exists {
		ordKeep = &orderKeep{
			waiting: make(chan *iot.Request),
		}
		rm.data[req.Key] = ordKeep
	}


	go func() {
		ordKeep.waiting <- req
	}()
	rm.nextExecution(req.Key)
}

func (rm *requestMap) removeRequest(req *iot.Request) {
	rm.locker.Lock()
	rm.locker.Unlock()

	ordKeep, exists := rm.data[req.Key]
	if !exists {
		return
	}

}

func (rm *requestMap) nextExecution(key string) {
	ordKeep, exists := rm.data[key]
	if !exists {
		return
	}

	if ordKeep.next == nil {

	}

	if ordKeep.writeInProgress {
		return
	}

	if ordKeep.next.IsRead() {
		ordKeep.readsInProgress += 1
	} else if ordKeep.readsInProgress == 0 {
		ordKeep.writeInProgress = true
	}
	log.Printf("Info: Starting execution for request: %s\n", ordKeep.next.ID)
	ventager.Fire(configs.Events.RequestExecutionBegin, ordKeep.next)
	ordKeep.next = nil

	rm.nextExecution(key)
}

func (rm *requestMap) getNextRequest(key string) {
	ordKeep, exists := rm.data[key]
	if !exists {
		return
	}


}