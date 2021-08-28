/*
	Given 2 arrays, create a function that let's a user knwow (true/false)
	whether these two arrays contain any common items
	For example:
	  array1 := [4]string{"a", "b", "c", "x"}
	  array2 := [4]string{"z", "y", "i"}
	  should return false
	  ------------
	  array1 = [4]string{"a", "b", "c", "x"}
	  array2 = [4]string{"z", "y", "x"}
	  should return true

	2 parameters - arrays - no size limit
	return true or false

	o(n^2)
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num = 1000000
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	array1  [num]string
	array2  [num]string
)

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		array1[i] = randSeq(5)
		array2[i] = randSeq(5)
	}
}

func containsCommonItem(array1, array2 [num]string) bool {
	now := time.Now()
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				fmt.Println("containsCommonItem1 took time: ",
					time.Now().Sub(now).Milliseconds(),
					"ms",
				)
				return true
			}
		}
	}

	fmt.Println("containsCommonItem1 took time: ",
		time.Now().Sub(now).Milliseconds(),
		"ms",
	)
	return false
}

func containsCommonItem2(array1, array2 [num]string) bool {
	// loop through first array and create object
	// where properties === items in the array
	now := time.Now()
	myMap := make(map[string]bool, len(array1))
	for i := 0; i < len(array1); i++ {
		if !myMap[array1[i]] {
			item := array1[i]
			myMap[item] = true
		}
	}

	// loop through second array and check if item
	// in ssecond array exists on created object.
	for j := 0; j < len(array2); j++ {
		if myMap[array2[j]] {
			fmt.Println("containsCommonItem2 took time: ",
				time.Now().Sub(now).Milliseconds(),
				"ms",
			)
			return true
		}
	}
	fmt.Println("containsCommonItem2 took time: ",
		time.Now().Sub(now).Milliseconds(),
		"ms",
	)
	return false
}

func containsCommonItem3(array1, array2 [num]string) bool {
	now := time.Now()
	myMap := make(map[string]bool, len(array1))
	for i := 0; i < len(array1); i++ {
		if !myMap[array1[i]] {
			item := array1[i]
			myMap[item] = true
		}
	}

	for _, arr := range array2 {
		_, ok := myMap[arr]
		if ok {
			fmt.Println("containsCommonItem3 took time: ",
				time.Now().Sub(now).Milliseconds(),
				"ms",
			)
			return true
		}
	}

	fmt.Println("containsCommonItem3 took time: ",
		time.Now().Sub(now).Milliseconds(),
		"ms",
	)
	return false
}

// o(a + b)

func main() {
	containsCommonItem(array1, array2)

	containsCommonItem2(array1, array2)

	containsCommonItem3(array1, array2)
}
