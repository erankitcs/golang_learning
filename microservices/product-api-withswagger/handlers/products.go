package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProduct(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)
	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id

}
