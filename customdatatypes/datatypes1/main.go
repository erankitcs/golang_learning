package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/customdatatypes/datatypes1/organisation"
)

func main() {
	fmt.Println("Custom Type Demo 1")
	p := organisation.NewPerson("Ankit", "Singh")
	err := p.SetTwitterHandler("@ankit63")
	if err != nil {
		fmt.Printf("An error occurred while setting handler %s", err.Error())
	}

	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.ID())

}
