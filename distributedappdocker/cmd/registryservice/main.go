package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/erankitcs/golang_learning/distributedapp/registry"
)

func main() {
	registry.SetupRegistryService()
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
		fmt.Println("Registry Service Started. Press key- s key to stop.")
		var s string
		fmt.Scanln(&s)
		if s == "s" {
			fmt.Println("Stop requested by user")
			srv.Shutdown(ctx)
			cancle()

		}
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registry service.")

}
