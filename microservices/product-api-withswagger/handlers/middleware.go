package handlers

import (
	"context"
	"net/http"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
)

//Middleware Validate Product
func (p *Products) MiddlewareProductsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			http.Error(rw, "Unable to parse product request", http.StatusBadRequest)
			return
		}
		//validate the product
		errs := p.v.Validate(prod)

		if len(errs) != 0 {
			p.l.Println("[ERROR] Validating the product", errs)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			//http.Error(rw, fmt.Sprintf("Error validating the product: %s", err), http.StatusBadRequest)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
