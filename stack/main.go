package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	left := len(s.items) - 1

	if left == 0 {
		return -1
	}

	item, items := s.items[left], s.items[0:left]
	s.items = items
	return item
}

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
