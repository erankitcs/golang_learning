package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/erankitcs/golang_learning/distributedapp/log"
	"github.com/erankitcs/golang_learning/distributedapp/registry"
	"github.com/erankitcs/golang_learning/distributedapp/service"
	"github.com/erankitcs/golang_learning/distributedapp/teacherportal"
)

func main() {
	err := teacherportal.ImportTemplates()

	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	globalServiceURL := fmt.Sprintf("http://teacherportalservice:%v", port)
	var r registry.Registration
	r.ServiceName = registry.TeacherPortalService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = globalServiceURL + "/services"
	r.HeartbeatURL = globalServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		teacherportal.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal.")

}
