package main

import (
	"test-go/datastructure"
)

func main() {
	var myTree = datastructure.NewBinaryTree()
	myTree.AddNode(5)
	myTree.AddNode(10)
	myTree.AddNode(15)
	myTree.AddNode(3)
	myTree.AddNode(4)
	datastructure.PreOrder(myTree.Root)
}
