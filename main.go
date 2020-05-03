package main

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/grpcserver"
	"github.com/hisitra/hedron/src/restserver"
)

func main() {
	configs.Load()
	almanac.CreateBaseDir()
	go restserver.New().Start()
	grpcserver.New().Start()
}
