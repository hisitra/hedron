package communicator

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/internalclient"
)

func BulkSet(record []byte) chan []*comcn.Output {
	resChan := make(chan []*comcn.Output)
	go bSHandler(record, resChan)
	return resChan
}

func bSHandler(record []byte, resChan chan []*comcn.Output) {
	var resArr []*comcn.Output
	singleResChan := make(chan *comcn.Output)

	go func() {
		singleResChan <- setSelf(record)
	}()

	for _, fellow := range configs.Node.Fellows {
		go func(address string) {
			singleResChan <- internalclient.Set(address, record)
		}(fellow)
	}

	for i := 0; i < len(configs.Node.Fellows)+1; i++ {
		resArr = append(resArr, <-singleResChan)
	}
	close(singleResChan)
	resChan <- resArr
}

func setSelf(record []byte) *comcn.Output {
	return almanac.Set(record)
}