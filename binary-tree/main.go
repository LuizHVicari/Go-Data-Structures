package main

import "fmt"

type BinaryTree struct {
	root * Node
}

type Node struct {
	value int
	leftNode * Node
	rightNode * Node
}

func (t *BinaryTree) initTree(value int) {
	node := Node {value: value, leftNode: nil, rightNode: nil}
	t.root = &node
}

func (t *BinaryTree) insertNode(value int, root *Node) {
	if value < root.value {
		if root.leftNode != nil {
			t.insertNode(value, root.leftNode)
		} else {
			node := Node {value: value, leftNode: nil, rightNode: nil}
			root.leftNode = &node
		}
	} else if value > root.value {
		if root.rightNode != nil {
			t.insertNode(value, root.rightNode)
		} else {
			node := Node {value: value, leftNode: nil, rightNode: nil}
			root.rightNode = &node
		}
	}
}

func (t *BinaryTree) printTree(root Node){
	fmt.Printf("(%d) ", root.value)
	if root.leftNode != nil {
		t.printTree(*root.leftNode)
	}
	if root.rightNode != nil {
		t.printTree(*root.rightNode)
	}
}

func (t BinaryTree) countNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return t.countNodes(root.leftNode) + t.countNodes(root.rightNode) + 1
}

func main() {
	t := BinaryTree {}
	t.initTree(2)
	t.insertNode(1, t.root)
	t.insertNode(3, t.root)
	t.insertNode(4, t.root)
	t.insertNode(5, t.root)
	t.insertNode(-3, t.root)
	t.insertNode(-4, t.root)
	t.insertNode(-1, t.root)
	t.printTree(*t.root)
	fmt.Printf("Number of nodes: %d", t.countNodes(t.root))
}