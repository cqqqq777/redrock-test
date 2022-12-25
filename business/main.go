package main

import "redrock-test/business/services/entry"

func main() {
	entry.InitDataBase()
	entry.InitRouter()
}
