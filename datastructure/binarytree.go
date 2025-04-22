package datastructure

type TreeNode struct {
	info      int
	leftNode  *TreeNode
	rightNode *TreeNode
}

func NewTreeNode(info int) *TreeNode {
	return &TreeNode{info: info, leftNode: nil, rightNode: nil}
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{root: nil}
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *BinaryTree) AddNode(node TreeNode, info int) {
	if tree.IsEmpty() {
		tree.root = NewTreeNode(info)
		return
	}

}