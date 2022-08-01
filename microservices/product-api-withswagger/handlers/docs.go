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

import "github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation error message returned as a string
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A list of product return into the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteProduct listSingleProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// The product to be updated.
	// in: body
	// required: true
	Body data.Product
}
