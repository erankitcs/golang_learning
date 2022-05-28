package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/customdatatypes/datatypes2/organisation"
)

func main() {
	fmt.Println("Custom Type Demo 1")
	p := organisation.NewPerson("Ankit", "Singh")
	err := p.SetTwitterHandler("@ankit63")
	fmt.Printf("%T\n", organisation.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred while setting handler %s", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())

}
