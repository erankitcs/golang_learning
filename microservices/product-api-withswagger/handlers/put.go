package handlers

import (
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("[DEBUG] Updating product for id: %#v\n", prod.ID)
	err := data.UpdateProduct(prod)
	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}
}
