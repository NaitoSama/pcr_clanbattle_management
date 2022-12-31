package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type BossConf struct {
	One   BossValue `json:"one"`
	Two   BossValue `json:"two"`
	Three BossValue `json:"three"`
	Four  BossValue `json:"four"`
	Five  BossValue `json:"five"`
}

type BossValue struct {
	First  int64 `json:"first"`
	Second int64 `json:"second"`
	Third  int64 `json:"third"`
	Fourth int64 `json:"fourth"`
	Fifth  int64 `json:"fifth"`
}

func ReadJson() BossConf {
	jsonfile, err := os.Open("statics/data/boss.json")
	if err != nil {
		fmt.Println("fail to open file")
	}
	defer jsonfile.Close()
	decoder := json.NewDecoder(jsonfile)

	conf := BossConf{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return conf
}
