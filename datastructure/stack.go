package datastructure

type StackNode struct {
	info int
	next *StackNode
}

func NewStackNode(info int) *StackNode {
	return &StackNode{info: info, next: nil}
}

type StackList struct {
	top *StackNode
}

func NewStackList() *StackList {
	return &StackList{top: nil}
}

func (list *StackList) IsEmpty() bool {
	return list.top == nil
}

func (list *StackList) Clear() {
	list.top = nil
}

func (list *StackList) Push(info int) {
	if list.IsEmpty() {
		list.top = NewStackNode(info)
		return
	}
	newNode := NewStackNode(info)
	newNode.next = list.top
	list.top = newNode
}

func (list *StackList) Pop() int {
	if list.IsEmpty() {
		return -1
	}
	popValue := list.top.info
	list.top = list.top.next
	return popValue
}

func (list *StackList) Peek() int {
	return list.top.info
}
