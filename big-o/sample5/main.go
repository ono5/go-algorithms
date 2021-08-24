package main

import "fmt"

func booooo(n []int) {
	for _, _ = range n {
		fmt.Println("booooo!")
	}
}

func arrayOfHiNTimes(n int) []string {
	hiArray := []string{}
	for i := 0; i < n; i++ {
		hiArray = append(hiArray, "hi")
	}
	return hiArray
}

func main() {
	booooo([]int{1, 2, 3, 4, 5})
	fmt.Println(arrayOfHiNTimes(6))
}
