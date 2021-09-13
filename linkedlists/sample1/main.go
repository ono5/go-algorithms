package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value int) *LinkedList {
	head := &Node{
		value: value,
		next:  nil,
	}
	return &LinkedList{
		head:   head,
		tail:   head,
		length: 1,
	}
}

func (l *LinkedList) append(value int) {
	newNode := NewNode(value)
	l.tail.next = newNode
	l.tail = newNode
	l.length++
}

func (l *LinkedList) prepend(value int) {
	newNode := NewNode(value)
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *LinkedList) printList() []int {
	array := []int{}
	currentNode := l.head
	for {
		array = append(array, currentNode.value)
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}
	return array
}

func (l *LinkedList) insert(index, value int) {
	if index >= l.length {
		l.append(value)
		return
	}
	newNode := NewNode(value)
	leader := l.traverseToIndex(index - 1)
	holdingPointer := leader.next
	leader.next = newNode
	newNode.next = holdingPointer
	l.length++
}

func (l *LinkedList) traverseToIndex(index int) *Node {
	counter := 0
	currentNode := l.head
	for {
		if counter != index {
			break
		}
		currentNode = currentNode.next
		counter++
	}
	return currentNode
}

func (l *LinkedList) remove(index int) {
}

func main() {
	myLinkedList := NewLinkedList(10)
	myLinkedList.append(5)
	myLinkedList.append(16)
	myLinkedList.prepend(1)
	myLinkedList.insert(2, 99)
	myLinkedList.insert(1, 33)
	fmt.Println(myLinkedList.printList())
}
