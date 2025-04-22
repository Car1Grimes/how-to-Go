package main

import "fmt"

type Node struct {
	info int
	next *Node
}

func NewNode(info int) *Node {
	return &Node{info: info, next: nil}
}

type StackList struct {
	top *Node
}

func NewStackList() *StackList {
	return &StackList{top: nil}
}

func (list *StackList) isEmpty() bool {
	return list.top == nil
}

func (list *StackList) clear() {
	list.top = nil
}

func (list *StackList) push(info int) {
	if list.isEmpty() {
		list.top = NewNode(info)
		return
	}
	newNode := NewNode(info)
	newNode.next = list.top
	list.top = newNode
}

func (list *StackList) pop() int {
	if list.isEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	popValue := list.top.info
	list.top = list.top.next
	return popValue
}

func (list *StackList) peek() int {
	return list.top.info
}

func (list *StackList) traverseAndPrint() {
	if list.isEmpty() {
		fmt.Println("No items found")
		return
	}
	cur := list.top
	for cur != nil {
		fmt.Print(cur.info)
		fmt.Print(" ")
		cur = cur.next
	}
}

func main() {
	var myStack = NewStackList()
	myStack.push(7)
	myStack.push(5)
	myStack.push(3)
	myStack.push(1)
	myStack.pop()
	myStack.traverseAndPrint()
	fmt.Printf("\nTop item: %d", myStack.peek())
}
