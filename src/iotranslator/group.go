package iotranslator

import "errors"

type ReqGroup struct {
	Reqs map[string][]*Request
}

func NewReqGroup() *ReqGroup {
	return &ReqGroup{Reqs: map[string][]*Request{}}
}

func (g *ReqGroup) IsReqAddable(req *Request) bool {
	reqArr, exists := g.Reqs[req.Key]
	if !exists || len(reqArr) == 0 {
		// Only way a write request can enter a group.
		return true
	}

	// Since group is not empty for this key,
	// writes will not be accepted.
	if !req.IsRead() {
		return false
	}

	// AcceptING read only if first request is read.
	// So read requests can keep populating the array.
	return reqArr[0].IsRead()
}

func (g *ReqGroup) AddReq(req *Request) error {
	if !g.IsReqAddable(req) {
		return errors.New("request not addable")
	}
	_, exists := g.Reqs[req.Key]
	if !exists {
		g.Reqs[req.Key] = []*Request{}
	}
	g.Reqs[req.Key] = append(g.Reqs[req.Key], req)
	return nil
}

func (g *ReqGroup) GetAll() []*Request {
	var allReqs []*Request
	for _, reqArr := range g.Reqs {
		for _, req := range reqArr {
			if req.ExecutionStatus == 0 {
				req.ExecutionStatus = 1
				allReqs = append(allReqs, req)
			}
		}
	}
	return allReqs
}

func (g *ReqGroup) ExecutionComplete() bool {
	for _, reqArr := range g.Reqs {
		for _, req := range reqArr {
			if req.ExecutionStatus != 2 {
				return false
			}
		}
	}
	return true
}
