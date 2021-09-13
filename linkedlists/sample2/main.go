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

func (l *LinkedList) Push(value string) {
	node := NewNode(value)

	// ヘッダがからの場合は一番最初のNodeということ
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

type Node struct {
	value string
	next  *Node
	prev  *Node
}

func NewNode(value string) *Node {
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
	l.Push("A")
	l.Push("B")
	l.Push("C")
	l.Push("D")

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
