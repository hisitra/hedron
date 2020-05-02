package almanac

import (
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"log"
	"path"
)

func Get(key string) ([]byte, *comcn.Output) {
	contents := readFile(path.Join(configs.Storage.BaseLocation, key+".record"))
	if contents == nil {
		return nil, iot.NotFoundResponse("Record not found.")
	}
	return contents, iot.SuccessResponse("")
}

func Set(recordJSON []byte) *comcn.Output {
	rec, err := DecodeRecord(recordJSON)
	if err != nil {
		log.Println("Record decoding failed because:", err)
		return iot.BadRequestResponse("Bad record provided.")
	}

	err = writeFile(
		path.Join(configs.Storage.BaseLocation, rec.Key+".record"),
		recordJSON)

	if err != nil {
		return iot.InternalServerErrorResponse("Failed to write record.")
	}
	return iot.SuccessResponse("")
}
