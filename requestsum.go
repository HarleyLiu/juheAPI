package juheAPI

import (
	"time"
)

type RequestCounter struct {
	ids             map[string]uint8
	blackList       map[string]int64
	maxRequestCount uint8
	delayTime       int64
}

func NewRequestCounter(counter *map[string]uint8, bl *map[string]int64, maxRequestCount uint8, delayTime int64, beClear bool) *RequestCounter {
	if maxRequestCount == 0 || *counter == nil || len(*counter) != 0 || *bl == nil || len(*bl) != 0 {
		return nil
	}
	rc := &RequestCounter{
		ids:             *counter,
		blackList:       *bl,
		maxRequestCount: maxRequestCount,
		delayTime:       delayTime,
	}
	if beClear {
		go rc.clearBlackList()
	}
	return rc
}

func (rc *RequestCounter) IsMoreRequst(id string) (err error) {
	if id == "" {
		return ErrorNoRequestID
	}
	if _, ok := rc.blackList[id]; ok {
		return ErrorMoreRequest
	}
	count, ok := rc.ids[id]
	if !ok {
		rc.ids[id] = 1
		return
	}
	if count < rc.maxRequestCount {
		rc.ids[id]++
		return
	}
	rc.AddBlackList(id)
	return ErrorMoreRequest
}

func (rc *RequestCounter) AddBlackList(id string) {
	if id == "" {
		return
	}
	rc.blackList[id] = time.Now().Unix()
}

func (rc *RequestCounter) RemoveBlackList(id string) {
	if id == "" {
		return
	}
	delete(rc.blackList, id)
}

func (rc *RequestCounter) clearBlackList() {
	for {
		for k, v := range rc.blackList {
			if time.Now().Unix() > v+rc.delayTime {
				delete(rc.ids, k)
				delete(rc.blackList, k)
			}
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
