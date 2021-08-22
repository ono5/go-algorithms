package main

import "fmt"

var boxes = []string{"a", "b", "c", "d", "e"}

// Big-O O(n2) - Quadratic Time
func logAllPairsOfArray(array []string) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Println(array[i], array[j])
		}
	}
}

// O(a + b) O(a*b)
func compressBoxesTwice(boxes, boxes2 []string) {
	for _, b := range boxes {
		fmt.Println(b)
	}

	for _, b := range boxes2 {
		fmt.Println(b)
	}
}

func main() {
	logAllPairsOfArray(boxes)

	compressBoxesTwice(boxes, boxes)
}
