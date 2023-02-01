package main

import (
	"redrock-test/boot"
	"redrock-test/cron"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.DatabaseInit()
	cron.Cron()
	boot.InitRouters()
}
