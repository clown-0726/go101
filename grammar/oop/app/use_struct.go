package app

import "fmt"

// 对于 OOM 来说，go 只支持封装，不支持继承和多态

type Node struct {
	value int
	left  *Node
	right *Node
}

// 工厂函数，用来实现 Node 的构造
// 在 go 中无法判断这个 Node 是创建在堆上还是在栈上，由编译器决定
func CreateNode(value int) *Node {
	return &Node{value: value}
}

func (node *Node) Print() {
	fmt.Println(node.value)
}

// 接受者函数，也是结构体的方法，接受者可以安全的接受 nil
func (node *Node) SetValue(value int) {
	node.value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}

func UseNode() {
	/*
			                  root:3
		        node:2                     root:1
		node:nil      node:nil       node:0      node:8
	*/

	//var root Node // 这时候 root 为 nil
	root := Node{value: 3}
	root.left = &Node{value: 2}
	root.right = &Node{value: 1, left: nil, right: nil}
	root.right.left = new(Node) // 这时候 new 出来的 Node 的 value 是默认初始值 0
	root.right.right = CreateNode(9)
	root.right.right.SetValue(8)

	// 遍历
	root.Traverse()
}

func UseNodeMany() {
	nodes := []Node{
		{3, nil, nil},
		{},
		{2, &Node{4, nil, nil}, nil},
	}
	fmt.Println(nodes)
}
