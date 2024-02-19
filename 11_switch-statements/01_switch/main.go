package main

import "fmt"

func switchFriends(name string) {
	// golang中一旦switch命中就会直接退出，不会继续往下执行
	// 如果想要往下执行需要使用fallthrough
	switch name {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}

}

func main() {

	switchFriends("Daniel")
	switchFriends("Medhi")
	switchFriends("Jenny")
	switchFriends("Sushant")
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
