package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/erankitcs/golang_learning/distributedapp/log"
	"github.com/erankitcs/golang_learning/distributedapp/registry"
	"github.com/erankitcs/golang_learning/distributedapp/service"
)

func main() {
	log.Run("./app.log")
	var r registry.Registration
	hostname, port := "logservice", "4000"
	serviceURL := fmt.Sprintf("http://%v:%v", hostname, port)
	globalServiceURL := fmt.Sprintf("http://logservice:%v", port)
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceURL
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = globalServiceURL + "/services"
	r.HeartbeatURL = globalServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		r,
		hostname,
		port,
		log.RegisterHandler,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
