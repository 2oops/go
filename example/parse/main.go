package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

	// 使用new()关键字创建指针
	str := new(string)
	*str = "2oops"
	fmt.Println(*str)
}