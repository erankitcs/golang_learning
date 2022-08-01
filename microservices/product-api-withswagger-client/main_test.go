package main

import (
	"testing"

	"github.com/erankitcs/golang_learning/microservices/productapiclient/client"
	"github.com/erankitcs/golang_learning/microservices/productapiclient/client/products"
)

func TestProductAPIClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	param := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(param)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(prods.Payload)
}
