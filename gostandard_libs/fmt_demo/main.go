package main

import (
	"fmt"
)

func main() {
	var firstNumber int
	println("Whats you first number ?")
	fmt.Scanln("%d", &firstNumber)

	var secondNumber int
	println("Whats you second number ?")
	fmt.Scanln("%d", &secondNumber)

	fmt.Printf("Two number sum is %v", firstNumber+secondNumber)

}
