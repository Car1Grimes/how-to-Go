package datastructure

type QueueNode struct {
	info int
	next *QueueNode
}

func NewQueueNode(info int) *QueueNode {
	return &QueueNode{info: info, next: nil}
}

type QueueList struct {
	head *QueueNode
	tail *QueueNode
}

func NewQueueList() *QueueList {
	return &QueueList{head: nil, tail: nil}
}

func (queue *QueueList) IsEmpty() bool {
	return queue.head == nil
}

func (queue *QueueList) Clear() {
	queue.head = nil
	queue.tail = nil
}

func (queue *QueueList) Enqueue(info int) {

	newNode := NewQueueNode(info)
	if queue.IsEmpty() {
		queue.head = newNode
		queue.tail = newNode
		return
	}
	queue.tail.next = newNode
	queue.tail = newNode
}

func (queue *QueueList) Dequeue() int {
	if queue.IsEmpty() {
		return -1
	}
	value := queue.head.info
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
	return value
}

func (queue *QueueList) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.head.info
}
