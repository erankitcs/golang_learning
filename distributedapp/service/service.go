package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/erankitcs/golang_learning/distributedapp/registry"
)

func Start(ctx context.Context, reg registry.Registration, host, port string,
	registerHandlerFun func()) (context.Context, error) {
	registerHandlerFun()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Please press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
