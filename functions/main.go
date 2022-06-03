package main

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/erankitcs/golang_learning/functions/simplemaths"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func ReadSomething() error {
	r := BadReader{errors.New("my nonsense error")}
	_, err := r.Read([]byte("Reading Something..."))
	if err != nil {
		fmt.Printf("an erro occurred .. %s", err)
		return err
	}
	return nil

}

//More error handling
type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrophic occurred in the reader")

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	/*if sr.count == 2 {
		//panic(errors.New("another error"))
		panic(errCatastrophicReader)
	} */
	if sr.count > 3 {
		return 0, io.EOF
	}
	sr.count += 1
	return sr.count, nil
}

func (sr *SimpleReader) Close() error {
	println("closing reader.")
	return nil
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is ok")
		} else if p != nil {
			panic("an unexpected error occurred and we do not want to recover")
		}
	}()
	defer func() {
		println("before for loop.")
	}()
	for {
		value, readerErr := r.Read([]byte("text reading..."))

		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println(value)
	}

	defer func() {
		println("after for-loop")
	}()

	return nil
}

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
	powerFun := powerOfTwo()
	powerVal := powerFun()
	println(powerVal)
	powerVal = powerFun()
	println(powerVal)
	// Bad state of Annonymous function
	var funcs []func() int64
	for i := 0; i < 10; i++ {
		/// here before function is being called i value is changed, hence we changed to cleanI
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	for _, f := range funcs {
		println(f())
	}

	// Error Handling
	ReadSomething()
	//
	println("\nDeep dive error handling.")
	if err := ReadFullFile(); err != nil {
		fmt.Printf("Something bad occured: %s", err)
	}
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

// Statefull function

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}
