package main

import "fmt"

// Create a function that reverses a string
// 'Hi My Name is Andrei' should be:
// 'ierdnA si eman yM iH'

func reverse(str string) string {
	// 文字列が3文字以上を対象とする
	// if str == "" || len(str) < 2 {
	// 	return "there are no propery strings"
	// }

	backwords := ""
	totalItems := len(str) - 1
	for i := totalItems; i >= 0; i-- {
		backwords += string(str[i])
	}
	return backwords
}

func main() {
	keyword := "Hi, My name is Selfnote!"
	fmt.Println(reverse(keyword))
}
