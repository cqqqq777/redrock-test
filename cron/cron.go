package cron

import (
	"github.com/robfig/cron/v3"
)

func Cron() {
	c := cron.New()
	_, err := c.AddFunc("@every 1h", RedisToMysql)
	if err != nil {
		panic(err)
	}
	c.Start()
}

func RedisToMysql() {

}
