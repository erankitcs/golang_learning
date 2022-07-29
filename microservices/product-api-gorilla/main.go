package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erankitcs/golang_learning/microservices/product-api-gorilla/handlers"
	"github.com/gorilla/mux"
)

var bindAddress = ":9090"

func main() {
	//Create a logger
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	// New Product handler
	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)

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
