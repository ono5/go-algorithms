package main

import (
	"log"
	"time"
)

var snacks = []string{
	"sherbet", "ice", "donuts", "caramel", "candy",
	"gummy", "crepe", "sable", "gelato", "cookie",
}

func findCookie(array []string) {
	now := time.Now()
	for _, a := range array {
		if a == "cookie" {
			log.Println("Found Cookie!")
		}
	}
	log.Println("Call to find Cooki took", time.Now().Sub(now))
}

func main() {
	findCookie(snacks)
}
