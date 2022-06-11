package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/erankitcs/golang_learning/distributedapp/grades"
	"github.com/erankitcs/golang_learning/distributedapp/log"
	"github.com/erankitcs/golang_learning/distributedapp/registry"
	"github.com/erankitcs/golang_learning/distributedapp/service"
)

func main() {
	var r registry.Registration
	hostname, port := "localhost", "6000"
	serviceURL := fmt.Sprintf("http://%v:%v", hostname, port)
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceURL
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	fmt.Println(r)
	ctx, err := service.Start(
		context.Background(),
		r,
		hostname,
		port,
		grades.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down Grading service.")
}
