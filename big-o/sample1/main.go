package main

import (
	"log"
	"time"
)

var snacks = []string{
	"sherbet", "ice", "donuts", "caramel", "candy",
	"gummy", "crepe", "sable", "gelato", "cookie",
}

// Big-O = o(n) linear time
func findCookie(array []string) {
	now := time.Now()
	for _, a := range array {
		if a == "cookie" {
			log.Println("Running")
			log.Println("Found Cookie!")
			break
		}
	}
	log.Println("Call to find Cooki took", time.Now().Sub(now))
}

var boxes = []int{0, 1, 2, 3, 4, 5}

// Big-O = o(1) constant time
func logFirstTowBoxes(boxes []int) {
	log.Println(boxes[0])
	log.Println(boxes[1])
}

func main() {
	// logFirstTowBoxes(boxes)
	findCookie(snacks)
}
