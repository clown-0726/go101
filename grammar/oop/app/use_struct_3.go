package app

// 通过嵌入的方式来扩充类型

type MyNodeEmbedding struct {
	*Node
}

func (myNodeEmbedding *MyNodeEmbedding) postOrderEmbedding() {
	if myNodeEmbedding == nil || myNodeEmbedding.Node == nil {
		return
	}

	left := MyNodeEmbedding{myNodeEmbedding.left}
	left.postOrderEmbedding()
	right := MyNodeEmbedding{myNodeEmbedding.right}
	right.postOrderEmbedding()
	myNodeEmbedding.Print()
}

func PostOrderTraverseEmbedding() {
	root := Node{value: 3}
	root.left = &Node{value: 2}
	root.right = &Node{value: 1, left: nil, right: nil}
	root.right.left = new(Node) // 这时候 new 出来的 Node 的 value 是默认初始值 0
	root.right.right = CreateNode(9)
	root.right.right.SetValue(8)

	myNodeEmbedding := MyNodeEmbedding{&root}
	myNodeEmbedding.postOrderEmbedding()
}
