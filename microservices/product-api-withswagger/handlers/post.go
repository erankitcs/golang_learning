package handlers

import (
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
