package juheAPI

import (
	"time"
)

const (
	delayTime = 10 //second
)

var (
	ids             map[string]uint8
	blackList       map[string]int64
	maxRequestCount uint8
)

func init() {
	ids = make(map[string]uint8)
	blackList = make(map[string]int64)
	go clearBlackList()
}

func isMoreRequst(id string) (err error) {
	if id == "" {
		return ErrorNoRequestID
	}
	if _, ok := blackList[id]; ok {
		return ErrorMoreRequest
	}
	count, ok := ids[id]
	if !ok {
		ids[id] = 1
		return
	}
	if count < maxRequestCount {
		ids[id]++
		return
	}
	addBlackList(id)
	return ErrorMoreRequest
}

func addBlackList(id string) {
	if id == "" {
		return
	}
	blackList[id] = time.Now().Unix()
}

func clearBlackList() {
	for {
		for k, v := range blackList {
			if time.Now().Unix() > v+delayTime {
				delete(ids, k)
				delete(blackList, k)
			}
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
