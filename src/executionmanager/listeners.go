package executionmanager

import (
	"github.com/hisitra/hedron/src/configs"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"github.com/hisitra/ventager"
	"log"
)

// TODO: Update Ventager for ListenAll and ListenOne methods.
// ListenOne channels are deleted after a single fire.

func Start() {
	go func() {
		for {
			listenArrivals()
		}
	}()
	go func() {
		for {
			listenCompletions()
		}
	}()
}

func listenArrivals() {
	newReq, ok := (<-ventager.Listen(configs.Events.RequestArrival)).(*iot.Request)
	if !ok {
		log.Printf(
			"Warning: Execution Manager received invalid data from %s event.\n",
			configs.Events.RequestArrival)
		return
	}
	log.Printf("Info: New request arrived with ID: '%s', Key: '%s' and Mode: '%s'\n", newReq.ID, newReq.Key, newReq.Mode)
	reqMap.acceptRequest(newReq)
}

func listenCompletions() {
	completedReq, ok := (<-ventager.Listen(configs.Events.RequestExecuted)).(*iot.Request)
	if !ok {
		log.Printf(
			"Warning: Execution Manager received invalid data from %s event.\n",
			configs.Events.RequestExecuted)
		return
	}
	reqMap.removeRequest(completedReq)
}
