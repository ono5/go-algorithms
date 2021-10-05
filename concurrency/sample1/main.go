package main

import "time"

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyCh <-chan string, donChan chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			donChan <- true
		}
	}
}
func main() {
	helloCh := make(chan string, 1)
	goodbyCh := make(chan string, 1)
	doneChan := make(chan bool)

	go receiver(helloCh, goodbyCh, doneChan)
	go sendString(helloCh, "hello!")
	go sendString(goodbyCh, "goodbye!")

	// doneが受信されるまでブロック
	<-doneChan
}
