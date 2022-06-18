package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/erankitcs/golang_learning/grpcdemo/server/pb/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	return nil, nil
}

func (s *employeeServer) GetAll(req *messages.GetAllRequest, stream messages.EmployeeService_GetAllServer) error {
	return nil
}

func (s *employeeServer) Save(ctx context.Context, req *messages.EmployeeRequest) (*messages.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeServer) SaveAll(stream messages.EmployeeService_SaveAllServer) error {
	return nil
}

func (s *employeeServer) AddPhoto(stream messages.EmployeeService_AddPhotoServer) error {
	return nil
}
