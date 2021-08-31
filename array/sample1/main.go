package main

import "fmt"

type myArray struct {
	length int
	data   map[int]string
}

func NewMyArray() myArray {
	return myArray{
		length: 0,
		data:   make(map[int]string),
	}
}

func (m *myArray) get(index int) string {
	return m.data[index]
}

func (m *myArray) push(item string) int {
	m.data[m.length] = item
	m.length += 1
	return m.length
}

func (m *myArray) pop() string {
	lastItem := m.data[m.length-1]
	delete(m.data, m.length-1)
	m.length--
	return lastItem
}

func (m *myArray) delete(index int) string {
	item := m.data[index]
	m.shiftItems(index)
	return item
}

func (m *myArray) shiftItems(index int) {
	for i := index; i < m.length-1; i++ {
		m.data[i] = m.data[i+1]
	}
	delete(m.data, m.length-1)
	m.length--
}

func main() {
	newArray := NewMyArray()
	newArray.push("hi")
	newArray.push("you")
	newArray.push("!")
	// newArray.pop()
	// newArray.pop()
	newArray.delete(0)
	newArray.push("are")
	newArray.push("nice")

	newArray.delete(1)

	fmt.Println(newArray)
}
