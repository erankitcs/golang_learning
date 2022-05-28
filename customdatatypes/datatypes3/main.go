package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/customdatatypes/datatypes3/organisation"
)

func main() {
	fmt.Println("Custom Type Demo 3")
	p := organisation.NewPerson("Ankit", "Singh", organisation.NewEuropeanUnionIdentifier("123-45-33", "UK"))
	err := p.SetTwitterHandler("@ankit63")
	fmt.Printf("%T\n", organisation.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred while setting handler %s", err.Error())
	}
	p.First = "NewName"
	println(p.First)
	println(p.Name.First)
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())
	println(p.Country())

}
