package main

import "fmt"

type Queue struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (q *Queue) initializeQueue(value int) {
	node := Node{next: nil, value: value}
	q.head = &node
}

func (q *Queue) enqueue(value int) {
	node := Node{value: value, next: nil}
	headNode := q.head
	for headNode.next != nil {
		headNode = headNode.next
	}
	headNode.next = &node
}

func (q *Queue) dequeue() Node {
	node := q.head
	q.head = node.next
	return *node
}

func (q *Queue) printQueue() {
	node := q.head
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
}

func main() {
	q := Queue{}
	q.initializeQueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.dequeue()
	q.printQueue()
}