package messages

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v", name)
}
