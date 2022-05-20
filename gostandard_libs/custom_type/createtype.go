package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	empId     int
	firstName string
	lastName  string
}

func main() {
	emps := make([]employee, 3)
	emps = append(emps, employee{1, "Ankit", "Singh"})
	emps = append(emps, employee{2, "Nikhil", "Karkara"})
	emps = append(emps, employee{3, "Anuj", "Agr"})

	eType := reflect.TypeOf(emps)
	fmt.Printf("Emp Type %s \n", eType)
	newEmpList := reflect.MakeSlice(eType, 0, 0)
	newEmpList = reflect.Append(newEmpList, reflect.ValueOf(employee{4, "NewEmp", "here"}))
	fmt.Printf("First List %v \n", emps)
	fmt.Printf("New List %v \n", newEmpList)

}
