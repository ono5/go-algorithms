package main

import "fmt"

// Google Question
// Given an array = [2,5,1,2,3,5,1,2,4]
// It should return 2

// Given an array = [1,5,1,2,3,5,1,2,4]
// It should return 1

// Given an array = [2,3,4,5]
// It should return 0

func slicesearch(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return input[i]
			}
		}
	}
	return 0
}

func hashSearch(input []int) int {
	myMap := make(map[int]int)
	for i := 0; i < len(input); i++ {
		_, ok := myMap[input[i]]
		if ok {
			return input[i]
		} else {
			myMap[input[i]] = i
		}
	}
	return 0
}

func main() {
	input := []int{2, 5, 1, 5, 3, 5, 1, 2, 4}
	fmt.Println(slicesearch(input))
	input = []int{1, 5, 1, 2, 3, 5, 1, 2, 4}
	fmt.Println(slicesearch(input))
	input = []int{1}
	fmt.Println(slicesearch(input))

	fmt.Println("-----------------------------")

	input = []int{2, 5, 1, 5, 3, 5, 1, 2, 4}
	fmt.Println(hashSearch(input))
	input = []int{1, 5, 1, 2, 3, 5, 1, 2, 4}
	fmt.Println(hashSearch(input))
	input = []int{1}
	fmt.Println(hashSearch(input))

}
