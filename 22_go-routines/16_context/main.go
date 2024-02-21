package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("outer worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
	// 没有指定标签的 break 只会跳出 switch/select 语句；
	// 若不能使用 return 语句跳出的话，可为 break 跳出标签指定的代码块：

LOOP:
	for {
		fmt.Println("inner worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("获取到退出信号")
			return
		default:
			userID := ctx.Value("userID")
			fmt.Println("用户ID", userID)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(1)
	// go worker(ctx)
	// time.Sleep(time.Second * 3)
	// cancel() // 通知子goroutine结束
	// wg.Wait()
	// fmt.Println("over")

	ctx2, cancel := context.WithCancel(context.Background())
	// context传值
	wg.Add(1)
	valCtx := context.WithValue(ctx2, "userID", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	// wg.Wait()

	time.Sleep(3 * time.Second)
	cancel()

	// 等待协程处理完毕才退出函数
	wg.Wait()
}
