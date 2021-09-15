package main

import "fmt"

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if n.value >= value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}

	fmt.Println("value", n.value)
	printNode(n.left)
	printNode(n.right)
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	if n.value >= value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func main() {
	t := &Tree{}
	t.insert(2).
		insert(7).
		insert(5).
		insert(2).
		insert(6).
		insert(5).
		insert(11).
		insert(5).
		insert(9).
		insert(4)

	printNode(t.node)

	fmt.Println("Is there 0?: ", t.node.exists(0))
	fmt.Println("Is there 9?", t.node.exists(9))
	fmt.Println("Is there 11?", t.node.exists(11))
	fmt.Println("Is there 99?", t.node.exists(99))
}
