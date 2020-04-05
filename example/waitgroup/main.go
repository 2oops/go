package main

import (
	"fmt"
	"sync"
	"net/http"
)

func main()  {
	// 等待组
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func (url string)  {
			// 函数完成时减1
			defer wg.Done()
			// 访问每个网址
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	// 当等待组计数器不等于 0 时阻塞直到变 0
	wg.Wait()
	fmt.Println("all over")
}