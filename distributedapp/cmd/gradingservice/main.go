package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/erankitcs/golang_learning/distributedapp/grades"
	"github.com/erankitcs/golang_learning/distributedapp/registry"
	"github.com/erankitcs/golang_learning/distributedapp/service"
)

func main() {
	var r registry.Registration
	hostname, port := "localhost", "6000"
	serviceURL := fmt.Sprintf("http://%v:%v", hostname, port)
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceURL
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
	<-ctx.Done()
	fmt.Println("Shutting down Grading service.")
}
