package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	username string
	nickname string
	sex      uint8
	birthday time.Time
}

// 结构体字段未大写，对外不可见
func main() {
	u := user{
		username: "坤坤",
		nickname: "阿坤",
		sex:      20,
		birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
