package main

import (
	"fmt"
	"main/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding埋め込み
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Node.Print()
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, new(tree.Node), new(tree.Node)}
	root.Right.Right = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
