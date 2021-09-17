package main

import (
	"fmt"
	"time"
)

func smaileWorker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(500 * time.Millisecond)
		out <- n
	}
}

func produce(in chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n", i)
		in <- i
		i++
	}
}

func consume(out <-chan int) {
	for n := range out {
		fmt.Printf("<- Recv job: %d\n", n)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	// Workerを4つ作る
	for i := 0; i < 4; i++ {
		go smaileWorker(in, out)
	}
	go produce(in)

	consume(out)
}
