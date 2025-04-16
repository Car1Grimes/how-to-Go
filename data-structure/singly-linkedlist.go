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

func IsEmpty(list *SinglyLinkedList) bool {
	return list.head == nil
}

func ClearList(list *SinglyLinkedList) {
	list.head = nil
	list.tail = nil
}

func AddNode(list *SinglyLinkedList, info int) {
	newNode := NewNode(info)
	if IsEmpty(list) {
		list.tail = newNode
		list.head = newNode
		return
	}
	list.tail.next = newNode
	list.tail = newNode
}

func ContainsNode(list *SinglyLinkedList, info int) bool {
	if IsEmpty(list) {
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

func DeleteLast(list *SinglyLinkedList) {
	if IsEmpty(list) {
		return
	}
	if list.head == list.tail {
		ClearList(list)
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

func TraverseAndPrint(list *SinglyLinkedList) {
	if IsEmpty(list) {
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
	AddNode(myList, 2)
	AddNode(myList, 5)
	AddNode(myList, 10)
	AddNode(myList, 6)
	TraverseAndPrint(myList)
	fmt.Println("Checking if list contains 6...")
	fmt.Println(ContainsNode(myList, 6))
	DeleteLast(myList)
	TraverseAndPrint(myList)
}
