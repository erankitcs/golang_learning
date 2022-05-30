package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/functions/simplemaths"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	fmt.Printf("Calling Add. \n")
	_ = simplemaths.Add(2, 3)
	total := simplemaths.Add(3, 10)
	fmt.Printf("Total : %f \n", total)
	fmt.Printf("Variadic function test.\n")
	inputs := []float64{12.3, 14.3, 19.0}
	total = simplemaths.Sum(inputs...)
	fmt.Printf("Total : %f \n", total)

	//Method Test.
	svObj := simplemaths.NewSemanticVersion(1, 1, 0)
	fmt.Printf("%s\n", svObj.String())
	fmt.Printf("%s\n", simplemaths.AnotherString(svObj))
	fmt.Printf("I am done.\n")

	//State Modification
	// Traditional way
	//svObj = svObj.IncrementMajor()
	//svObj = svObj.IncrementMinor()
	// Pointer based reciever.
	// p := &svObj (Not required. Go will take care of it.)
	// p.IncrementMajor()
	svObj.IncrementMajor()
	svObj.IncrementPatch()
	fmt.Printf("%s\n", svObj.String())

	// Annonymous function
	func() {
		println("My First Annonymous function")
	}()
	// Function as variable.
	a := func(name string) string {
		fmt.Printf("My Second Annonymous function by %s\n", name)
		return "hello from annonymous"
	}

	msg := a("Ankit")
	println(msg)

	// Function return from function.
	addFunction := mathExpression1()
	println(addFunction(2, 4))
	addExp := mathExpression(AddExpr)
	println(addExp(2, 4))
	subExp := mathExpression(SubtractExpr)
	println(subExp(2, 4))

	//Function as parameter
	fmt.Printf("%f\n", double(2, 3, mathExpression(AddExpr)))

	//Satefull function
}

// Function Return from Function. Returns a function of two input and one output.
func mathExpression1() func(float64, float64) float64 {
	return simplemaths.Add
}

func mathExpression(exp MathExpr) func(float64, float64) float64 {
	switch exp {
	case AddExpr:
		return simplemaths.Add
	case SubtractExpr:
		return simplemaths.Subtract
	case MultiplyExpr:
		return simplemaths.Multiply
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}

}

//Function as parameter.

func double(f1, f2 float64, mathExp func(float64, float64) float64) float64 {
	return 2 * mathExp(f1, f2)
}
