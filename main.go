package main

import (
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/grpcserver"
)

func main() {
	configs.Load()
	grpcserver.New().Start()
}
