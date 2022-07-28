package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/erankitcs/golang_learning/microservices/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		regx := regexp.MustCompile(`/([0-9]+)`)
		g := regx.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)
	// Encoder is much faster and memory efficient
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to parse product data", http.StatusInternalServerError)
	}
	//rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handling PUT request")
	newprod := &data.Product{}
	err := newprod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse product request", http.StatusBadRequest)
	}
	//p.l.Printf("New Product %#v", newprod)
	data.AddProduct(newprod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
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
