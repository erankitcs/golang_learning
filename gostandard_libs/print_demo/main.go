package main

import "fmt"

func main() {
	var age = 42

	var out, _ = fmt.Print("I am ", age, " Year old.\n")
	print("Bytes written - ", out)

	var name = "Jeremy"

	fmt.Printf("\nMy name is %s and I am %d years old\n", name, age)

	var pi float32 = 3.141592

	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("Pi is %2.2f\n", pi)

	test := fmt.Sprintf("|%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	print(test)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 98.999, 12.3456, 12.01)

	fmt.Printf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	fmt.Printf("|%-7s|%-7s|%-7d|\n", "a", "ab", 100)

	type point struct {
		x, y int
	}

	p := point{1, 2}
	fmt.Printf("%v\n", p)

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	newPerson := Person{"Ankit", "Singh", 42}

	fmt.Printf("%T\n", newPerson)

	var isCool = true
	fmt.Printf("Value is %t\n", isCool)
	fmt.Printf("Value is %T\n", isCool)

	fmt.Printf("%c\n", 3)
}
