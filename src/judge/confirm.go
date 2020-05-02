package judge

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	iot "github.com/hisitra/hedron/src/iotranslator"
)

func DecideConfirm(pushResult []*comcn.Output, pushedRecord []byte, req *iot.Request) ([]byte, *comcn.Output) {
	if !isMajorityPositive(pushResult) {
		return nil, iot.InternalServerErrorResponse("Failed to push to majority.")
	}

	recordData, err := almanac.DecodeRecord(pushedRecord)
	if err != nil {
		return nil, iot.InternalServerErrorResponse("Failed to decode record.")
	}

	if req.Type == "C" || req.Type == "U" {
		recordData.Committed = recordData.Uncommitted
		recordData.Uncommitted = ""
		recordData.DeleteStatus = 0
		recordData.Version += 1

		recordJSON, err := almanac.EncodeRecord(recordData)
		if err != nil {
			return nil, iot.InternalServerErrorResponse("Failed to encode record.")
		}
		return recordJSON, iot.SuccessResponse("")
	}
	if req.Type == "D" {
		recordData.DeleteStatus = 2
		recordData.Committed = ""
		recordData.Uncommitted = ""
		recordData.Version += 1

		recordJSON, err := almanac.EncodeRecord(recordData)
		if err != nil {
			return nil, iot.InternalServerErrorResponse("Failed to encode record.")
		}
		return recordJSON, iot.SuccessResponse("")
	}

	return nil, iot.InternalServerErrorResponse("Invalid request type.")
}

func CheckConfirm(confirmResult []*comcn.Output) *comcn.Output {
	if !isMajorityPositive(confirmResult) {
		return iot.InternalServerErrorResponse("Failed to confirm on majority.")
	}
	return iot.SuccessResponse("")
}

func isMajorityPositive(responses []*comcn.Output) bool {
	var errorCount int
	for _, resp := range responses {
		if resp.Code != 200 {
			errorCount++
		}
	}
	return errorCount < len(responses)-errorCount
}