package configs

import "github.com/hisitra/confine"

var Server = &grpcServer{}

func Load() {
	confine.LoadAll(map[string]interface{}{
		"./src/configs/json/grpcServer.json": Server,
	})
}
