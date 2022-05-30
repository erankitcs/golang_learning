package simplemaths

import "errors"

// Variadic function...slice of type as inputs. It can be only last parameter.
func Sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

// Naming return value.
func Devide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cant devide by 0")
	}
	answer = p1 / p2
	return
}
