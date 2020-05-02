package executionqueue

import (
	iot "github.com/hisitra/hedron/src/iotranslator"
	"log"
	"time"
)

var ArrivalChan = make(chan *iot.Request)

func init() {
	go requestListener()
}

func requestListener() {
	for {
		newRequest := <-ArrivalChan
		log.Println("New request arrived with ID:", newRequest.ID)
		queue.locker.Lock()
		queue.addRequest(newRequest)
		queue.locker.Unlock()
		log.Println("Request added to Execution Queue.")
		time.Sleep(10 * time.Nanosecond)
	}
}
