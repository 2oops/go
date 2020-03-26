package main

import (
	"fmt"
	"container/list"
	"unsafe"
	"os/user"
)

type weekday int

const (
	sunday weekday = iota
	monday
	tuesday
	wednesday
)

var team [3]int = [3]int{1,2}

func main()  {
	fmt.Println(monday)

	// 数组遍历
	team[2] = 5
	for k, v := range team {
		fmt.Println(k,v)
	}

	// 多维数组
	var array [4][2]int
	array = [4][2]int{{1,2}, {2,3}, {3,4},{4,5}}
	var value int = array[1][1]
	fmt.Println(value)
	
	// 切片
	slice1 := []int{1,2,3,4}
	slice2 := []int{5,4,3,2,1}
	copy(slice1, slice2)

	fmt.Println(slice1) // [5,4,3,2]

	// 删除指定元素
	str := []string{"a","b", "c", "d", "e"}
	index := 2
	str = append(str[:index], str[index + 1:]...)
	fmt.Println(str) // [a b d e]

	// range
	slice := []int{1,2,3,4}
	for index, value := range slice {
		fmt.Println(index, value, &value, &slice[index])
	}
	for index := 2; index < len(slice); index++ {
		fmt.Println(slice[index])
	}

	// 多维切片
	slice3 := [][]int{{3}, {1,2,3}}
	slice3[0] = append(slice3[0], 4)
	fmt.Println(slice3)

	// map
	var mapLit map[string]int
	mapLit = map[string]int{"name": 1, "value":2 }
	fmt.Println(mapLit["name"])

	delete(mapLit, "name")
	fmt.Println(mapLit) // map[value: 2]

	// 列表
	myList := list.New()
	myList.PushFront("2oops")
	myList.PushBack(20)
	// 保存一个句柄
	element := myList.PushBack("center")
	// 插入
	myList.InsertBefore("left", element)
	myList.InsertAfter("right", element)
	myList.Remove(element)

	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value) // 2oops 20 left center right
	}

	// nil
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", num) // 0x0

	var m map[int]bool
	fmt.Println( unsafe.Sizeof( m ) ) // 8
	var c chan string
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	u, _ := user.Current()
	fmt.Println(u.Username)
}