package main

import "fmt"

// 函数执行到if语句中直接返回，未执行接下来的defer语句
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
