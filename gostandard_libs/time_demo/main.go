package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()

	fmt.Printf("Today is %d/%d/%d", month, day, year)

	time.Sleep(1 * time.Second)
	elapsed := time.Since(t)
	fmt.Printf("\nElasped time is %s \n", elapsed)

	// Reference time:
	// Mon Jan 2 15:04:05 MST 2006
	fmt.Printf("Current Time %v \n", t.Format("15:04:05"))
	fmt.Printf("Current Date %v \n", t.Format("Monday 02, Jan-2006"))
	fmt.Printf("Today is %v \n", t.Format("Monday 02, Jan-2006 at 15:04:05"))
}
