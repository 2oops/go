package main

import (
	"fmt"
)

func main()  {
	a := true
	if a {
		fmt.Println(a)
	} else {
		fmt.Println("false")
	}

	OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				fmt.Println(i, j) 
				break OuterLoop
			}
			if j == 5 {
				fmt.Println(i,j)
				break OuterLoop
			}
		}
	}

	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	fmt.Println(step)

	var i int
	// for {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	i++
	// }
	for i <= 10 {
		i++
	}
	fmt.Println(i)

	// 九九乘法表
	for k := 1; k <= 9; k ++ {
		for j := 1; j <= k; j++ {
			fmt.Printf("%d * %d = %d |", k, j, k*j)
		}
		fmt.Println()
	}

	// 通道遍历
	// c := make(chan int)
	// go func() {
	// 	c <- 1
	// 	c <- 2
	// 	c <- 3
	// }()
	// for v := range c {
	// 	fmt.Println(v)
	// }

	var b = true
	switch b {
		case true : 
		fmt.Println("int")
		fallthrough
		case 1 == 1:
		fmt.Println("fallthrough")
	}

	for x := 0; x < 10; x++ {
		if x == 2 {
			goto breakHere
		}
		if x == 5 {
			goto breakHere
		}
	}
	// 手动返回，避免执行进入标签
	return
	breakHere:
	fmt.Println("ok")
}