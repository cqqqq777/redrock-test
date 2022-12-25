package main

import "fmt"

func main() {
	ch := make(chan struct{}, 1)
	go func() {
		fmt.Println("出现")
		ch <- struct{}{}
	}()
	<-ch
}
