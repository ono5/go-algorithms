package main

import "fmt"

// O(x^3)
func printAllNumbersThenALlPairSums(numbers []int) {
	fmt.Println("there are the numbers:")
	for _, i := range numbers {
		fmt.Println(i)
	}

	fmt.Println("and these are their nums:")
	for _, i := range numbers {
		for _, j := range numbers {
			fmt.Println(j + i)
		}
	}
}

func main() {
	printAllNumbersThenALlPairSums([]int{1, 2, 3, 4, 5})
}
