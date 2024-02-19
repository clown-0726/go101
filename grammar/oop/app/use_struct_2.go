package app

type MyNode struct {
	Node *Node
}

func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := MyNode{myNode.Node.left}
	left.postOrder()
	right := MyNode{myNode.Node.right}
	right.postOrder()
	myNode.Node.Print()
}

func PostOrder() {
	root := Node{value: 3}
	root.left = &Node{value: 2}
	root.right = &Node{value: 1, left: nil, right: nil}
	root.right.left = new(Node) // 这时候 new 出来的 Node 的 value 是默认初始值 0
	root.right.right = CreateNode(9)
	root.right.right.SetValue(8)

	myNode := MyNode{&root}
	myNode.postOrder()
}
