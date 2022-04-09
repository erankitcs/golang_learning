package main

import (
	"fmt"
	"time"
)

func pinger(ping <-chan string, pong chan<- string) {
	for m := range ping {
		printAndDelay(m)
		pong <- "PlayerA- pong"
	}
}

func ponger(pong <-chan string, ping chan<- string) {
	for m := range pong {
		printAndDelay(m)
		ping <- "PlayerB- ping"
	}
}

func printAndDelay(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("Starting pingpong game.")
	ping := make(chan string)
	pong := make(chan string)

	// Creating two player
	go pinger(ping, pong)
	go ponger(pong, ping)

	// starting the game.
	ping <- "ping"

	for {

	}
}
