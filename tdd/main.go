package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	// terminalからの入力値はstring型
	input1, _ := strconv.Atoi(os.Args[1])
	input2, _ := strconv.Atoi(os.Args[2])

	result := sum(input1, input2)
	fmt.Printf("The sum of %d and %d is %d.\n", input1, input2, result)
}
