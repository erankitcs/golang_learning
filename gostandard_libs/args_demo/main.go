package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Slicing to include first read which is nothing but file name
	// Args are list of string be default.
	args := os.Args[1:]
	//fmt.Println(args)
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: Dinner Total <Total Ammount> <Tip Percentage>")
		fmt.Println("args_demo 10 5")
	} else {
		if len(args) != 2 {
			fmt.Println("Please enter some inputs. Type /help for more details abt input.")
		} else {
			// Business Logic here.
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Your total bill is %.2f", calculateTotal(float32(mealTotal), float32(tipAmount)))
		}
	}
}

func calculateTotal(mealTotal float32, tipAmount float32) float32 {
	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))
	return totalPrice
}
