package handlers

import (
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

// swagger:route DELETE /products/{id} products deleteProducts
// Delete a product from database
// responses:
//	201: noContentResponse
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] deleting record id", id)
	err := data.DeleteProduct(id)

	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
