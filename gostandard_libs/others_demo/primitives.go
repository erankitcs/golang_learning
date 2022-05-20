package main

import (
	"fmt"
	"reflect"
)

func main() {
	var amount int32

	amount = 32
	typename := reflect.TypeOf(amount).Name()
	fmt.Printf("The name of the type is %v \n", typename)
	fmt.Printf("The type is %v\n", reflect.TypeOf(amount))
	fmt.Printf("The kind is %v\n", reflect.TypeOf(amount).Kind())
	fmt.Printf("The value is %v\n", reflect.ValueOf(amount))

	newValue := GetReflectedValue(reflect.TypeOf(amount))
	fmt.Print(newValue)
	//test := reflect.New(newValue)
}

func GetReflectedValue(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}
