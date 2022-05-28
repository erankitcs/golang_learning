package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/customdatatypes/datatypes4/organisation"
)

type Name struct {
	First  string
	Last   string
	Middle []string
}

type OtherName struct {
	First string
	Last  string
}

func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

func main() {
	fmt.Println("Custom Type Demo 4")
	p := organisation.NewPerson("Ankit", "Singh", organisation.NewEuropeanUnionIdentifier("123-45-33", "UK"))
	err := p.SetTwitterHandler("@ankit63")
	fmt.Printf("%T\n", organisation.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred while setting handler %s", err.Error())
	}

	name1 := Name{
		First: "",
		Last:  "",
	}

	if name1.Equals(Name{}) {
		println("We match")
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
