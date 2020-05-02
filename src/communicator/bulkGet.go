package communicator

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/internalclient"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"time"
)

func BulkGet(req *iot.Request) chan []*comcn.Output {
	resChan := make(chan []*comcn.Output)
	go bGHandler(req, resChan)
	return resChan
}

func bGHandler(req *iot.Request, resChan chan []*comcn.Output) {
	var resArr []*comcn.Output
	singleResChan := make(chan *comcn.Output)

	go func() {
		singleResChan <- getSelf(req.Key)
	}()

	for _, fellow := range configs.Node.Fellows {
		go func(address string) {
			for {
				res := internalclient.Get(address, req)
				if res.Message != iot.EarlierRequestFoundResponse().Message {
					singleResChan <- res
					break
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(fellow)
	}

	for i := 0; i < len(configs.Node.Fellows)+1; i++ {
		resArr = append(resArr, <-singleResChan)
	}
	close(singleResChan)
	resChan <- resArr
}

func getSelf(key string) *comcn.Output {
	record, res := almanac.Get(key)
	if res.Code != 200 {
		return res
	}
	return iot.NewResponse(200, "Success", record)
}