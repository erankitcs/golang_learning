package main

import (
	"fmt"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func main() {
	var firstNumber int
	println("Whats you first number ?")
	fmt.Scanln("%d", &firstNumber)

	var secondNumber int
	println("Whats you second number ?")
	fmt.Scanln("%d", &secondNumber)

	fmt.Printf("Two number sum is %v", firstNumber+secondNumber)

	fmt.Println("\nMessage Formating program.")

	fileName := "test.txt"

	showMessage(INFO, fmt.Sprintf("About to open %s", fileName))

	file, err := os.Open("test.txt")
	if err != nil {
		showMessage(ERROR, err.Error())
	}
	defer file.Close()

}

func showMessage(messageType messageType, message string) {
	switch messageType {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}
