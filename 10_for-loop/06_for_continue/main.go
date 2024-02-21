package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			// continue ,如果是偶数跳过本次循环
			continue
		}
		fmt.Printf("非偶数:%d \n", i)
		if i >= 50 {
			break
		}
	}
}
