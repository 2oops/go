package main

import (
	"fmt"
	// "time"
	// "runtime"
	// "sync"
	"bytes"
)

func joinString(list ...string) string {
  var b bytes.Buffer // 定义一个字节缓冲, 快速地连接字符串
  for _, str := range list {
    b.WriteString(str)
  }
  return b.String()
}

func fibonacci(n int) (res int)  {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return res
}

func main() {
	fmt.Println(joinString("aa","bb", "cc")) // aabbcc

	result := 0
	for i := 0; i < 10; i ++ {
		result = fibonacci(i)
		fmt.Println(i, result)
	}

	// go func() {
	// 	var times int
	// 	for {
	// 		times++
	// 		fmt.Println("tick", times)
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	
	// var input string
	// fmt.Scanln(&input)

	// 循环接收通道数据
	// ch := make(chan int)
	// go func() {
	// 	for i := 5; i > 2; i -- {
	// 		ch <- i
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	// for data := range ch {
	// 	fmt.Println(data)
	// 	// 如果不终止会发现宕机或报错
	// 	if data == 3 {
	// 		break
	// 	}
	// }

	// 超时机制
	// ch := make(chan int)
	// quit := make(chan bool)
	// go func()  {
	// 	for {
	// 		select {
	// 		case num := <-ch: 
	// 			fmt.Println("num:", num)
	// 		case <- time.After(3 * time.Second):
	// 			fmt.Println("overtime")
	// 			quit <- true
	// 		}
	// 	}
	// }()
	// for i := 0; i < 5; i++ {
	// 	ch <- i
	// 	time.Sleep(time.Second)
	// }
	// <- quit
	// fmt.Println("end")

	// cpu核数
	// cpuNum := runtime.NumCPU()
	// fmt.Println("core num:", cpuNum)

	// // 读写互斥锁
	// var (
	// 	count int 
	// 	countGuard sync.RWMutex
	// )
	// func getCount() int {
	// 	countGuard.RLock()
	// 	//函数退出时解除锁定
	// 	defer countGuard.RUnlock
	// 	return count
	// }
}