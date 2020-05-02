package executioner

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/communicator"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"github.com/hisitra/hedron/src/judge"
	"log"
)

func ExecuteAll(reqs []*iot.Request) {
	for _, req := range reqs {
		go execute(req)
	}
}

func execute(req *iot.Request) {
	log.Println("Executing request:", req.ID)

	expiry := getExpiryNotifier(req)

	var pullResult []*comcn.Output
	select {
	case <-expiry:
		log.Println("Request:", req.ID, "Expired.")
		req.SendResponse(iot.TimeoutResponse(""))
		return
	case pr := <-communicator.BulkGet(req):
		log.Println("Pull result arrived for request:", req.ID)
		pullResult = pr
	}

	pushRecord, res := judge.DecidePush(pullResult, req)
	if res.Code != 200 {
		req.SendResponse(res)
		return
	}
	if req.IsRead() {
		log.Println("Sending response for request:", req.ID)
		rec, err := almanac.DecodeRecord(pushRecord)
		if err != nil {
			log.Println("Failed to decode record for read request:", req.ID)
			req.SendResponse(iot.InternalServerErrorResponse("Failed to decode record."))
			return
		}
		req.SendResponse(iot.NewResponse(200, "Success.", []byte(rec.Committed)))
		return
	}

	var pushResult []*comcn.Output
	select {
	case <-expiry:
		log.Println("Request:", req.ID, "Expired.")
		req.SendResponse(iot.TimeoutResponse(""))
		return
	case pr := <-communicator.BulkSet(pushRecord):
		log.Println("Push result arrived for request:", req.ID)
		pushResult = pr
	}

	confirmedRecord, res := judge.DecideConfirm(pushResult, pushRecord, req)
	if res.Code != 200 {
		req.SendResponse(res)
		return
	}

	var confirmResult []*comcn.Output
	select {
	case <-expiry:
		log.Println("Request:", req.ID, "Expired.")
		req.SendResponse(iot.TimeoutResponse("Request timed out. Unconfirmed write"))
		return
	case cr := <-communicator.BulkSet(confirmedRecord):
		log.Println("Confirmation result arrived for request:", req.ID)
		confirmResult = cr
	}

	res = judge.CheckConfirm(confirmResult)
	if res.Code != 200 {
		req.SendResponse(res)
		return
	}

	log.Println("Request:", req.ID, "execution successful. Sending response...")
	req.SendResponse(iot.NewResponse(200, "success", nil))
	log.Println("Sent response for request:", req.ID)
}
