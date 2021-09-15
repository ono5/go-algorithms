package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func process2(ch chan<- int) {
	// 0 - 3000のランダム数値を生成
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Microsecond)
	ch <- n
}

func main() {
	ch := make(chan int)
	go process2(ch)

	fmt.Println("Waiting for process")
	n := <-ch
	fmt.Printf("Process took %dms\n", n)
}
