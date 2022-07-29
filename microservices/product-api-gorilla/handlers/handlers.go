// Package Product API.
//
// The purpose of this API is to provide list of products, add new product and update the existing product.
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: er.ankit.cs@gmail.com
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erankitcs/golang_learning/microservices/product-api-gorilla/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to parse product data", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handling PUT request")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unxpected Id format", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product update failed", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p Products) MiddlewareProductsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to parse product request", http.StatusBadRequest)
			return
		}
		//validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Validating the product", err)
			http.Error(rw, fmt.Sprintf("Error validating the product: %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
