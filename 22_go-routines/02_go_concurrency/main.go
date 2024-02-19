package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
	// 如果不做等待，此处会直接返回。
	time.Sleep(5 * time.Second)
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
