package main

import (
	"fmt"
	"reflect"
)

func main() {

	type person struct {
		personId  int
		firstName string
		lastName  string
	}

	newPerson := person{0, "Ankit", "Singh"}

	fmt.Printf("My person is %s %s with Id %d\n", newPerson.firstName, newPerson.lastName, newPerson.personId)
	fmt.Printf("Type is %v \n", reflect.TypeOf(newPerson))
	fmt.Printf("Value is %v \n", reflect.ValueOf(newPerson))
	fmt.Printf("Kind of type is %v \n", reflect.ValueOf(newPerson).Kind())
	type employee struct {
		empId     int
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmp := employee{0, "Ankit", "Singh"}
	newCus := customer{100, "Ankit", "Singh", "NAB"}

	addPerson(newEmp)
	addPerson(newCus)

}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "employee":
			empString := "insert into emp (?,?,?)"
			fmt.Printf("SQL String %s \n", empString)
			fmt.Printf("Added %v \n", v.Field(1))
		case "customer":
			custString := "insert into cust (?,?,?)"
			fmt.Printf("SQL String %s \n", custString)
			fmt.Printf("Added %v \n", v.Field(1))
		}
		return true
	} else {
		return false
	}
}
