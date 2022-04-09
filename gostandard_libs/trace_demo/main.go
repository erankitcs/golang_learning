package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	AddRandomNumbers()
}

func AddRandomNumbers() {

	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)

	time.Sleep(2 * time.Second)

	var result = firstNumber * secondNumber

	fmt.Printf("Result of 2 numbers is %d\n", result)
}
