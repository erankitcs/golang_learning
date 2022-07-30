package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/data"
	"github.com/erankitcs/golang_learning/microservices/product-api-withswagger/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var bindAddress = ":9090"

func main() {
	//Create a logger
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	v := data.NewValidation()
	// New Product handler
	ph := handlers.NewProduct(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareProductsValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareProductsValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	//creating new server
	s := http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	//Starting the server
	go func() {
		l.Printf("Starting Server on Port %s\n", bindAddress)
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error in starting the server: %s\n", err)
			os.Exit(1)
		}
	}()

	//Gracefull shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	//block util a signal recieved
	sig := <-c
	log.Println("Got Signal: ", sig)
	ctx, cancelFun := context.WithTimeout(context.Background(), 30*time.Second)
	cancelFun()
	s.Shutdown(ctx)

}
