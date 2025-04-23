package main

import (
	"fmt"
	"test-go/datastructure"
)

func main() {
	var myTree = datastructure.NewBinaryTree()
	myTree.AddNode(5)
	myTree.AddNode(10)
	myTree.AddNode(12)
	myTree.AddNode(3)
	myTree.AddNode(15)
	var existsNode = datastructure.PreOrderSearch(myTree.Root, 15)
	if !existsNode {
		fmt.Println("Node not found")
		return
	} 
		fmt.Println(existsNode)
}
