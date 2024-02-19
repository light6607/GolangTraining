package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		// 1.加了fallthrough后，会直接运行【紧跟的后一个】case或default语句，不论条件是否满足都会执行		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
