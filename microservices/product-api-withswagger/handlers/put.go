package handlers

import (
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation
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
