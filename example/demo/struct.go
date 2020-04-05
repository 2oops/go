package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Wheel struct {
	Size float32
}
type Engine struct {
	Power string
	Type string
}
type Car struct {
	Wheel
	Engine
}

func main()  {
	fmt.Println("aa")
	
	p := Person{
		"2oops",
		20,
		"male", // 要逗号
	}
	fmt.Println(p) // {2oops 20 male}

	c := Car {
		Wheel: Wheel {
			Size: 18,
		},
		Engine: Engine {
			Power: "14T",
			Type: "ppp",
		},
	}
	fmt.Println(c)
}