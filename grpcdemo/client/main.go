package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/erankitcs/golang_learning/grpcdemo/client/pb/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const serverAddress = ":9000"

type TlsType string

const (
	mTlsConn = TlsType("mtls")
	tlsConn  = TlsType("tls")
)

func loadTLSCredentials(tlstype TlsType) (credentials.TransportCredentials, error) {
	serverCACert, err := ioutil.ReadFile("../cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCACert) {
		return nil, fmt.Errorf("not able to load Server CA certificate")
	}

	config := &tls.Config{
		RootCAs:    certPool,
		ServerName: "gogrpcserver.com",
	}

	return credentials.NewTLS(config), nil
}

func main() {
	option := flag.Int("o", 1, "Command to run")
	enableTLS := flag.Bool("tls", false, "enable TLS for connection")
	enablemTLS := flag.Bool("mtls", false, "enable mTLS for connection")
	flag.Parse()
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	if *enableTLS {
		log.Printf("Connecting server %v with TLS=%v", serverAddress, *enableTLS)
		tlsCreds, err := loadTLSCredentials(tlsConn)
		if err != nil {
			log.Fatal("cant load Server CA Certificate for TLS")
		}
		transportOption = grpc.WithTransportCredentials(tlsCreds)
	}

	if *enablemTLS {
		log.Printf("Connecting server %v with TLS=%v", serverAddress, *enablemTLS)
		tlsCreds, err := loadTLSCredentials(mTlsConn)
		if err != nil {
			log.Fatal("cant load Client Certificate for mTLS")
		}
		transportOption = grpc.WithTransportCredentials(tlsCreds)
	}

	opt := []grpc.DialOption{
		transportOption,
	}

	conn, err := grpc.Dial(serverAddress, opt...)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer conn.Close()
	client := messages.NewEmployeeServiceClient(conn)
	switch *option {
	case 1:
		SendMetadate(client)
	}
}

func SendMetadate(client messages.EmployeeServiceClient) {
	fmt.Println("Sending metadata..")
	md := metadata.MD{}
	md["caller"] = []string{"Ankit"}
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	emp, err := client.GetByBadgeNumber(ctx, &messages.GetByBadgeNumberRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emp)
}
