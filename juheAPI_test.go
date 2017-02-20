package juheAPI

import (
	"log"
	"testing"
	"time"
)

func Test_Cehck(t *testing.T) {
	counter := make(map[string]uint8)
	bl := make(map[string]int64)
	rc := NewRequestCounter(&counter, &bl, 3, 60)
	icc, err := New("http://op.juhe.cn/idcard/query", "eb90badasd8c1162312319203190231284", rc)
	if err != nil {
		log.Println("init icc obj fail", err)
	}
	pass, err := icc.Check(GET, "13912345678", "410728198116134523", "赵卫娟")
	if err != nil {
		log.Println("check id card fail", err)
	} else if pass {
		log.Println("check pass")
	} else {
		log.Println("check not pass")
	}
	log.Println("test end ...")
}

func Test_MoreRequest(t *testing.T) {
	counter := make(map[string]uint8)
	bl := make(map[string]int64)
	rc := NewRequestCounter(&counter, &bl, 3, 60)
	icc, err := New("http://op.juhe.cn/idcard/query", "eb90badasd8c1162312319203190231284", rc)
	if err != nil {
		log.Println("init icc obj fail", err)
	}

	for i := 0; i < 10; i++ {
		_, err := icc.Check(GET, "13912345678", "410728198116134523", "赵卫娟")
		log.Println("request", i+1, err)
	}
}

func Test_MoreRequestWithDelay(t *testing.T) {
	counter := make(map[string]uint8)
	bl := make(map[string]int64)
	rc := NewRequestCounter(&counter, &bl, 3, 60)
	icc, err := New("http://op.juhe.cn/idcard/query", "eb90badasd8c1162312319203190231284", rc)
	if err != nil {
		log.Println("init icc obj fail", err)
	}

	for i := 0; i < 100; i++ {
		_, err := icc.Check(GET, "13912345678", "410728198116134523", "赵卫娟")
		log.Println("request", i+1, err)
		time.Sleep(1 * time.Second)
	}
}
