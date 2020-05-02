package executioner

import (
	iot "github.com/hisitra/hedron/src/iotranslator"
	"time"
)

func getExpiryNotifier(req *iot.Request) chan struct{} {
	expiryChan := make(chan struct{})
	go func() {
		for !req.IsExpired() {
			time.Sleep(time.Millisecond)
		}
		close(expiryChan)
	}()
	return expiryChan
}
