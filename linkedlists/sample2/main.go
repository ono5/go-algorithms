package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) First() *Node {
	return l.head
}

func (l *LinkedList) Last() *Node {
	return l.tail
}

func (l *LinkedList) Push(value int) {
	node := NewNode(value)

	// First Node?
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := NewLinkedList()
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

	fmt.Println("-------------")

	n = l.Last()
	for {
		println(n.value)
		n = n.Prev()
		if n == nil {
			break
		}
	}
}
