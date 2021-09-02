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

	fmt.Println("----------push---------")
	fmt.Println(newArray.push("hi"))
	fmt.Println(newArray.push("you"))
	fmt.Println(newArray.push("!"))
	fmt.Println(newArray.push("?"))

	fmt.Println("----------get---------")
	fmt.Println(newArray.get(0))
	fmt.Println(newArray.get(1))
	fmt.Println(newArray.get(2))
	fmt.Println(newArray.get(3))

	fmt.Println("----------pop---------")
	fmt.Println(newArray.pop())
	fmt.Println(newArray.pop())

	fmt.Println("---------delete--------")
	fmt.Println(newArray.delete(0))

	fmt.Println("----------Done---------")
	newArray.push("are")
	newArray.push("nice")

	newArray.delete(1)

	fmt.Println(newArray)
}
