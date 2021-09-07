package main

import "fmt"

type HashTable struct {
	data []int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([]int, size),
	}
}

func (h HashTable) _hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int([]rune(string(key[i]))[0]))
	}

	return hash
}

func main() {
	myHashTable := NewHashTable(50)
	fmt.Println(myHashTable._hash("Selfnote is the most wonderful blog in the world! but, there are many blogs is more wonderfule thant selfnote after all!"))
}
