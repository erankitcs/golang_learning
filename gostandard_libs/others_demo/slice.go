package main

import (
	"fmt"
	"reflect"
)

func main() {

	v := reflect.TypeOf(123)
	fmt.Println(reflect.SliceOf(v))
}
