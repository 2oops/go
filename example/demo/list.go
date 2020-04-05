package main
 
import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

func showNode(n *Node)  {
	for n != nil {
		fmt.Println(*n)
		n = n.next
	}
}

func main()  {
	// var head = new(Node)
	// head.data = 1
	// var node1 = new(Node)
	// node1.data = 2

	// head.next = node1
	// var node2 = new(Node)
	// node2.data = 3

	// node1.next = node2
	// showNode(node1)

	// 头插法
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head // tail记录头节点的位置，刚开始tail指针指向头节点
	// for i := 1; i < 10; i++ {
	// 	var node = Node{data: i}
	// 	node.next = tail // 新插入的node指向头节点
	// 	tail = &node // 重新赋值头节点
	// }
	// showNode(tail)

	// 尾插法
	for i := 20; i < 30; i ++ {
		var node = Node{data: i}
		(*tail).next = &node
		tail = &node
	}
	showNode(head)
	
	fmt.Println("args:", os.Args)
}