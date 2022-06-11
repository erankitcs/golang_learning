package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/erankitcs/golang_learning/distributedapp/registry"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancle()
	}()

	go func() {
		fmt.Println("Registry Service Started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancle()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registry service.")

}
