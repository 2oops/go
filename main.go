package main

import(
	"log"
	"net/http"
	_ "net/http/pprof"
	"fmt" 
	"math/rand"
	"time"
)

// 1. 通道
// 生产者每秒钟生成一个字符串，并通过通道传递给消费者，生产者通过两个goroutine并发运行
// 消费者在main()函数的goroutine进行处理

// 生产者无限的生产数据
func producer(header string, channel chan<- string)  {
	for {
		// 不停地往信道写入数据
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待一秒
		time.Sleep(time.Second)
	}
}

// 消费者
func customer(channel <- chan string)  {
	for {
		// 从信道中不断取出数据
		message := <- channel

		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	// 消费
	customer(channel)

	go func()  {
		log.Println(http.ListenAndServe("localhost:10000", nil))
	}()
}


