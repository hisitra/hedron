package judge

import (
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	iot "github.com/hisitra/hedron/src/iotranslator"
)

func DecidePush(pullResult []*comcn.Output, req *iot.Request) ([]byte, *comcn.Output) {
	if isMajorityDown(pullResult) {
		return nil, iot.InternalServerErrorResponse("Majority of nodes are down.")
	}

	records := res2Records(pullResult)
	highestVRec := highestVersionedRecord(records)
	if highestVRec == nil || highestVRec.Version == 0 || highestVRec.DeleteStatus == 2 {
		if req.Type == "C" {
			return newRecord(req.Key, req.Value)
		}
		return nil, iot.NotFoundResponse("Record not found.")
	}

	if req.Type == "C" {
		return nil, iot.NewResponse(409, "Key already exists.", nil)
	}
	if req.Type == "U" {
		highestVRec.Uncommitted = req.Value
	} else if req.Type == "D" {
		highestVRec.DeleteStatus = 1
	}

	recJSON, err := almanac.EncodeRecord(highestVRec)
	if err != nil {
		return nil, iot.InternalServerErrorResponse(err.Error())
	}
	return recJSON, iot.SuccessResponse("")
}

func isMajorityDown(responses []*comcn.Output) bool {
	var internalErrorCount int
	for _, resp := range responses {
		if resp.Code >= 500 {
			internalErrorCount++
		}
	}
	return internalErrorCount >= len(responses)-internalErrorCount
}

func res2Records(responses []*comcn.Output) []*almanac.Record {
	var records []*almanac.Record
	for _, res := range responses {
		if res.Code != 200 {
			continue
		}
		rec, err := almanac.DecodeRecord(res.Data)
		if err != nil {
			continue
		}
		records = append(records, rec)
	}
	return records
}

func highestVersionedRecord(records []*almanac.Record) *almanac.Record {
	// Here we are getting the "first" highest versioned record.
	// In reality there will be multiple records having the highest
	// version. Should we collect them all and then perform a check
	// if they are all same? because two records with same version
	// MUST have the same values. If not then it's a SYSTEM FAILURE.

	var highest *almanac.Record
	for _, rec := range records {
		if highest == nil || highest.Version < rec.Version {
			highest = rec
		}
	}
	return highest
}

func newRecord(key string, value string) ([]byte, *comcn.Output) {
	newRec, err := almanac.NewRecord(key, value)
	if err != nil {
		return nil, iot.BadRequestResponse(err.Error())
	}
	newRecJSON, err := almanac.EncodeRecord(newRec)
	if err != nil {
		return nil, iot.InternalServerErrorResponse("Failed to encode record.")
	}
	return newRecJSON, iot.SuccessResponse("")
}


