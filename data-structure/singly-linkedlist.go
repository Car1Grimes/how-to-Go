// vim: set ts=2 sw=2 expandtab:
package main

import "fmt"

type Node struct {
	info int
	next *Node
}

func NewNode(info int) *Node {
	return &Node{
		info: info,
		next: nil}
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{head: nil, tail: nil}
}

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *SinglyLinkedList) ClearList() {
	list.head = nil
	list.tail = nil
}

func (list *SinglyLinkedList) AddNode(info int) {
	newNode := NewNode(info)
	if list.IsEmpty() {
		list.tail = newNode
		list.head = newNode
		return
	}
	list.tail.next = newNode
	list.tail = newNode
}

func (list *SinglyLinkedList) ContainsNode(info int) bool {
	if list.IsEmpty() {
		return false
	}
	current := list.head
	for current != nil {
		if current.info == info {
			return true
		}
		current = current.next
	}
	return false
}

func (list *SinglyLinkedList) DeleteLast() {
	if list.IsEmpty() {
		return
	}
	if list.head == list.tail {
		list.ClearList()
		return
	}
	current := list.head
	for current.next != nil {
		if current.next == list.tail {
			list.tail = current
			current.next = nil
			return
		}
		current = current.next
	}
}

func (list *SinglyLinkedList) TraverseAndPrint() {
	if list.IsEmpty() {
		fmt.Println("No items found")
		return
	}
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.info)
		current = current.next
	}
}

func main() {
	var myList = NewSinglyLinkedList()
	myList.AddNode(2)
	myList.AddNode(5)
	myList.AddNode(10)
	myList.AddNode(6)
	myList.TraverseAndPrint()
	fmt.Println("Checking if list contains 6...")
	fmt.Println(myList.ContainsNode(6))
	myList.DeleteLast()
	myList.TraverseAndPrint()
}
