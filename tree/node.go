package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node *treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored")
		return
	}
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, new(treeNode), new(treeNode)}
	root.right.right = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)

	root.traverse()
}
