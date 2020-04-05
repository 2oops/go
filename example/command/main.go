package main

import (
	"fmt"
)
// go:generate go run main.go
// go:generate go version

// 定义一个数据写入器接口
type DataWriter interface {
	WriteData(data interface{}) error
}
// 定义一个文件结构，用于实现数据写入器
type file struct {

}
// 实现写入器的方法
func (d *file) WriteData(data interface{}) error {
	fmt.Println(data)
	return nil
}

// 一个服务接口需要开启功能和日志功能
type Service interface {
	Start()
	Log(string)
}
// 日志结构
type Logger struct {

}
// 实现服务的日志方法
func (l *Logger) Log(s string) {
	fmt.Println(s)
}
// 游戏服务的结构
type GameService struct {
	Logger // 嵌入日志服务
}
// 游戏服务只能实现开启功能
func (g *GameService) Start() {
	fmt.Println("GameService Start")
}

// 自定义错误类型
type myError struct {
	num int
	desc string
}

func (m myError) Error() string {
	return fmt.Sprintf("error, not the number I want")
}

func jude(f int) (float64, error) {
	if f < 3 {
		return -1, myError{num: f}
	}
	return 0, nil
}

func main() {
	// 实例化file struct
	var f = new(file)
	// 声明一个DataWriter接口
	var writer DataWriter
	// 将接口赋值f,即*file类型
	writer = f
	// 使用接口进行数据写入
	writer.WriteData(2)

	// 实例化游戏服务并赋给s
	var s Service = new(GameService)
	// s拥有了两个功能
	s.Start()
	s.Log("gameservice string")

	var x interface{}
	x = "hello"
	value, ok := x.(int) 
	// value := x.(int) // panic
	fmt.Println(value, ok) // 0 false

	getType(x)

	result, err := jude(2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// 动物名称到实例的映射
	animals := map[string] interface{} {
		"bird": new(bird),
		"pig": new(pig),
	}
	for name, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)
		fmt.Println(name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}
}

func getType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

// 会走的动物接口
type Flyer interface {
	Fly()
}
// 会飞的动物接口
type Walker interface {
	Walk()
}
type bird struct {

}
type pig struct {

}
// 鸟会飞的接口
func (b *bird) Fly() {
	fmt.Println("bird fly")
}
// 猪会走的接口
func (p *pig) Walk() {
	fmt.Println("pig walk")
}
// 鸟也会走的接口
func (b *bird) Walk() {
	fmt.Println("bird walk")
}