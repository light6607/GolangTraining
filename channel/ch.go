package main

import (
	"fmt"
	"time"
)

func main() {

	// 无缓冲channel
	ch := make(chan string)
	go func() {
		fmt.Println("等待3秒后输出")
		time.Sleep(3 * time.Second)
		ch <- "go route123123"
	}()
	v := <-ch
	fmt.Println("接收到的值为", v)

	// 无缓冲channel
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	ch2 := make(chan string, 5)
	go func() {
		fmt.Println("等待3秒后输出")
		time.Sleep(3 * time.Second)
		ch2 <- "go route123123"
		ch2 <- "go route123123123"
		// wg.Done()
		// 写完之后记得关闭，否则会出现 fatal error: all goroutines are asleep - deadlock!
		close(ch2)
	}()
	// wg.Wait()
	for v2 := range ch2 {
		fmt.Println("接收到的值为", v2)
	}

}
