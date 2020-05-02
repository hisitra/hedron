package executionqueue

import "github.com/hisitra/hedron/src/executioner"

func init() {
	go func() {
		for {
			feedExecutioner()
		}
	}()
}

func feedExecutioner() {
	queue.locker.Lock()
	defer queue.locker.Unlock()

	if len(queue.groups) == 0 {
		return
	}
	if queue.groups[0].ExecutionComplete() {
		queue.groups = queue.groups[1:]
		return
	}
	reqs := queue.groups[0].GetAll()
	if reqs == nil {
		return
	}
	executioner.ExecuteAll(reqs)
}
