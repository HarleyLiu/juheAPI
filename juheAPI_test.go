package juheAPI

import (
	"log"
	"testing"
)

func Test_Cehck(t *testing.T) {
	icc, err := New("http://op.juhe.cn/idcard/query", "eb90badasd8c1162312319203190231284")
	if err != nil {
		log.Println("init icc obj fail", err)
	}
	pass, err := icc.Check(GET, "410728198116134523", "赵卫娟")
	if err != nil {
		log.Println("check id card fail", err)
	} else if pass {
		log.Println("check pass")
	} else {
		log.Println("check not pass")
	}
	log.Println("test end ...")
}
