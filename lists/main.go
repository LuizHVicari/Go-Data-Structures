package main

import (
	"fmt"
	"errors"
)

type List struct {
	head * Node
	len int
}

type Node struct {
	next * Node
	previous * Node
	value int
}


func (l *List) initList(value int) {
	firstNode := Node { next: nil, previous: nil, value: value}
	l.head = &firstNode
	l.len = 1
}

func (l *List) addNodeToEnd(value int) {
	node := l.head
	
	for node.next != nil {
		node = node.next
	}

	node.next = &Node {next: nil, previous: node, value: value }
	l.len ++
}

func (l *List) addNodeToPos(value, pos int) error{
	if pos < 0 {
		return errors.New("cannot add node to a nagative position")
	}
	node := l.head

	if pos == 0 {
		l.head = &Node{ previous: nil, next: node, value: value}
		node.previous = l.head
	}

	for i := 0; i < pos - 1 && node.next != nil ; i++ {
		node = node.next
	}

	if node.next != nil {
		node.next = &Node{next: node.next, previous: node, value: value}
		node.next.next.previous = node.next
	} else {
		node.next = &Node{next: nil, previous: node, value: value}
	}
		
	l.len ++
	return nil
}

func (l *List) removeNodeFromValue(value int) error {
	node := l.head

	if node.value == value {
		l.head = node.next
		node.next.previous = nil
		l.len --
		return nil
	}

	for node.next != nil && node.value != value {
		node = node.next
	}

	if node.value != value {
		return errors.New("value not found in list")
	}

	next_node := node.next
	previous_node := node.previous

	previous_node.next = next_node
	next_node.previous = previous_node

	l.len --

	return nil
}

func (l *List) removeNodeFromPosition(pos int) error {
	if pos < 0 {
		return errors.New("cannot remove a node from a negative position")
	}

	if pos > l.len - 1 {
		return errors.New("cannot remove item from an index that does not exist")
	}

	if pos == 0 {
		l.head = l.head.next
		l.head.previous = nil

		return nil
	}

	node := l.head

	for i := 0; i < pos && node.next != nil; i ++ {
		node = node.next
	}
	
	if node.next != nil {
		next_node := node.next
		previous_node := node.previous
		
		previous_node.next = next_node
		next_node.previous = previous_node
	} else {
		previous_node := node.previous
		previous_node.next = nil
	}

	l.len --

	return nil
}

func (l List) printList() {
	node := l.head

	for node != nil {
		fmt.Printf("%d", node.value)
		node = node.next
		if node != nil {
			fmt.Printf(" -> ")
		}
	}

	fmt.Println()
}

func main() {
	l := List {}
	l.initList(0)
	l.addNodeToEnd(1)
	l.addNodeToEnd(2)
	l.addNodeToEnd(3)
	l.addNodeToPos(4, 5)
	l.removeNodeFromValue(2)
	l.removeNodeFromPosition(2)
	l.printList()
	fmt.Printf("len: %d", l.len)
}