package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/erankitcs/golang_learning/grpcdemo/server/pb/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const port = ":9000"

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // for no client verification
	}
	return credentials.NewTLS(config), nil

}

func main() {
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal(err)
	}
	opt := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
	}
	grpcServer := grpc.NewServer(opt...)
	messages.RegisterEmployeeServiceServer(grpcServer, new(employeeServer))
	log.Println("Starting server on port " + port)
	log.Println(grpcServer.GetServiceInfo())
	reflection.Register(grpcServer)
	grpcServer.Serve(listner)
}

type employeeServer struct {
	messages.UnimplementedEmployeeServiceServer
}

func (s *employeeServer) GetByBadgeNumber(ctx context.Context, req *messages.GetByBadgeNumberRequest) (*messages.EmployeeResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata recieved: %v\n", md)
	}
	for _, emp := range employees {
		if req.BadgeNumber == emp.BadgeNumber {
			return &messages.EmployeeResponse{Employee: &emp}, nil
		}
	}
	return nil, errors.New("employee not found")
}

func (s *employeeServer) GetAll(req *messages.GetAllRequest, stream messages.EmployeeService_GetAllServer) error {
	for _, emp := range employees {
		stream.Send(&messages.EmployeeResponse{Employee: &emp})
	}
	return nil
}

func (s *employeeServer) Save(ctx context.Context, req *messages.EmployeeRequest) (*messages.EmployeeResponse, error) {
	employees = append(employees, *req.Employee)
	return &messages.EmployeeResponse{Employee: req.Employee}, nil
}

func (s *employeeServer) SaveAll(stream messages.EmployeeService_SaveAllServer) error {
	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		employees = append(employees, *emp.Employee)
		stream.Send(&messages.EmployeeResponse{Employee: emp.Employee})
	}

	for _, emp := range employees {
		fmt.Println(emp)
	}
	return nil
}

func (s *employeeServer) AddPhoto(stream messages.EmployeeService_AddPhotoServer) error {
	imgData := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("File recieved with length : %v\n", len(imgData))
			return stream.SendAndClose(&messages.AddPhotoResponse{IsOK: true})
		}
		if err != nil {
			return err
		}
		imgData = append(imgData, data.Data...)
	}
}
