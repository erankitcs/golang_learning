package handlers

import (
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
	//d, err := json.Marshal(lp)
	// Encoder is much faster and memory efficient
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to parse product data", http.StatusInternalServerError)
	}
	//rw.Write(d)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handling PUT request")
	newprod := &data.Product{}
	err := newprod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse product request", http.StatusBadRequest)
	}
	//p.l.Printf("New Product %#v", newprod)
	data.AddProduct(newprod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unxpected Id format", http.StatusBadRequest)
	}
	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse product request", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product update failed", http.StatusInternalServerError)
		return
	}

}
