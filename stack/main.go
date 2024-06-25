package main

import "fmt"

type Stack struct {
	top *Node
}

type Node struct {
	value int
	next *Node
}

func (s *Stack) initializeStack(value int) {
	node := Node{value: value, next: nil}
	s.top = &node
}

func (s *Stack) push(value int) {
	newNode := Node{value: value, next: s.top}
	s.top = &newNode
}

func (s *Stack) pop() Node {
	node := s.top
	s.top = node.next
	return *node
}

func (s *Stack) printStack() {
	node := s.top

	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
}

func main() {
	s := Stack{}
	s.initializeStack(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.printStack()
	s.pop()
	s.printStack()
}