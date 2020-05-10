package configs

import (
	"github.com/hisitra/confine"
	"log"
)

var Server = &server{}
var Node = &node{}
var Storage = &storage{}
var Events = &events{}

func Load() {
	log.Println("Info: Loading configs...")
	confine.LoadAll(map[string]interface{}{
		"./src/configs/json/server.json":  Server,
		"./src/configs/json/node.json":    Node,
		"./src/configs/json/storage.json": Storage,
	})

	Events.RequestArrival = "request-arrival"
	Events.RequestExecutionBegin = "request-execution-begin"
	Events.RequestExecuted = "request-executed"

	log.Println("Info: Configs loaded successfully.")
}
