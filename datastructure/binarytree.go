package datastructure

import "fmt"

type TreeNode struct {
	Info      int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func NewTreeNode(info int) *TreeNode {
	return &TreeNode{Info: info, LeftNode: nil, RightNode: nil}
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.Root == nil
}

func addNode(root *TreeNode, info int) *TreeNode {
	if root == nil {
		return NewTreeNode(info)
	}
	if root.Info > info {
		root.LeftNode = addNode(root.LeftNode, info)
	}
	if root.Info < info {
		root.RightNode = addNode(root.RightNode, info)
	}
	return root
}

func (tree *BinaryTree) AddNode(info int) {
	tree.Root = addNode(tree.Root, info)
}

func visitNode(node *TreeNode) {
	fmt.Print(node.Info, " ")
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	visitNode(root)
	PreOrder(root.LeftNode)
	PreOrder(root.RightNode)
}
