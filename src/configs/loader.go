package configs

import (
	"github.com/hisitra/confine"
	"log"
)

var Server = &server{}
var Node = &node{}
var Storage = &storage{}

func Load() {
	log.Println("Loading configs...")
	confine.LoadAll(map[string]interface{}{
		"./src/configs/json/server.json":  Server,
		"./src/configs/json/node.json":    Node,
		"./src/configs/json/storage.json": Storage,
	})
	log.Println("Configs loaded successfully.")
}
