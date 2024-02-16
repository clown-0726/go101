package app

type MyTreeNode struct {
	TreeNode *Node
}

func (myTreeNode *MyTreeNode) postOrder() {
	if myTreeNode == nil || myTreeNode.TreeNode == nil {
		return
	}

	left := MyTreeNode{myTreeNode.TreeNode.left}
	left.postOrder()
	right := MyTreeNode{myTreeNode.TreeNode.right}
	right.postOrder()
	myTreeNode.TreeNode.Print()
}

func PostOrder() {
	root := Node{value: 3}
	root.left = &Node{value: 2}
	root.right = &Node{value: 1, left: nil, right: nil}
	root.right.left = new(Node) // 这时候 new 出来的 Node 的 value 是默认初始值 0
	root.right.right = CreateNode(9)
	root.right.right.SetValue(8)

	myTreeNode := MyTreeNode{&root}
	myTreeNode.postOrder()
}
