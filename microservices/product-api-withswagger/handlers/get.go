package handlers

import (
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

// swagger:route GET /products products listProducts
// Return a list of product.
// responses:
//	200: productsResponse
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	prods := data.GetProducts()
	rw.Header().Add("Content-Type", "application/json")
	err := data.ToJSON(prods, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a single  product.
// responses:
//	200: productResponse
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get single requested product")
	pid := getProductID(r)
	prod, err := data.GetProductByID(pid)
	switch err {
	case nil:
	case data.ErrorProductNotFound:
		p.l.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
