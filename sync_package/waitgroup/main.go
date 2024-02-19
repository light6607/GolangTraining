package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com/",
	"https://www.baidu.com",
	"https://youtube.com",
}

func headURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	fmt.Print(url, ":  \n", resp.Status)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	for _, url := range urls {
		fmt.Printf("访问域名:%s \n", url)
		go headURL(url, &wg)
	}
	wg.Wait()
}
