package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	ch := make(chan bool)
	go Work("goroutine1", ch)
	ch <- true
	go Work("goroutine2", ch)
	ch <- true
	go Work("goroutine3", ch)
	ch <- true
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string, ch chan bool) {
	time.Sleep(time.Second)
	select {
	case <-ch:
		fmt.Println(workName)
	} // 模拟逻辑处理

	wg.Done()
}
