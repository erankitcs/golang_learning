package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/erankitcs/golang_learning/grpcdemo/client/pb/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

//const serverAddress = ":9000"

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

	var clientCert tls.Certificate

	if tlstype == mTlsConn {
		// Load client's certificate and private key
		clientCert, err = tls.LoadX509KeyPair("../cert/client-cert.pem", "../cert/client-key.pem")
		if err != nil {
			return nil, err
		}
	}

	config := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
		//ServerName: "client.gogrpcserver.com",
	}

	return credentials.NewTLS(config), nil
}

func main() {
	option := flag.Int("o", 1, "Command to run")
	enableTLS := flag.Bool("tls", false, "enable TLS for connection")
	enablemTLS := flag.Bool("mtls", false, "enable mTLS for connection")
	serverhost := flag.String("serverhost", "", "Server Host for connection")
	serverport := flag.String("serverport", "9000", "Server Port for connection")
	flag.Parse()
	serverAddress := fmt.Sprintf("%s:%s", *serverhost, *serverport)
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
		log.Printf("Connecting server %v with mTLS=%v", serverAddress, *enablemTLS)
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
	case 2:
		GetByBadgeNumber(client)
	case 3:
		GetAll(client)
	case 4:
		SaveAll(client)
	case 5:
		AddPhoto(client)
	}
}

func SendMetadate(client messages.EmployeeServiceClient) {
	fmt.Println("Sending metadata..")
	md := metadata.MD{}
	md["caller"] = []string{"Ankit"}
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err := client.GetByBadgeNumber(ctx, &messages.GetByBadgeNumberRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(emp)
}

func GetByBadgeNumber(client messages.EmployeeServiceClient) {
	fmt.Println("Getting Employee by Badge Number")
	ctx := context.Background()
	res, err := client.GetByBadgeNumber(ctx, &messages.GetByBadgeNumberRequest{
		BadgeNumber: 2080,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Employee)

}

func GetAll(client messages.EmployeeServiceClient) {
	fmt.Println("Getting all Employees...")
	ctx := context.Background()
	stream, err := client.GetAll(ctx, &messages.GetAllRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Employee)
	}
}

func SaveAll(client messages.EmployeeServiceClient) {
	fmt.Println("Saving all employees...")
	saveEmployees := []messages.Employee{
		{
			FirstName:   "Ankit",
			LastName:    "Singh",
			BadgeNumber: 2008,
		},
		{
			FirstName:   "AnotherAnkit",
			LastName:    "Singh",
			BadgeNumber: 2020,
		}}
	ctx := context.Background()
	stream, err := client.SaveAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	doneCh := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				doneCh <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Employee)
		}
	}()
	for _, emp := range saveEmployees {
		err := stream.Send(&messages.EmployeeRequest{Employee: &emp})
		if err != nil {
			log.Fatal(err)
		}
	}
	stream.CloseSend()
	<-doneCh
}

func AddPhoto(client messages.EmployeeServiceClient) {
	fmt.Println("Adding photo..")
	f, err := os.Open("Penguins.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	md := metadata.MD{}
	md.Append("badgeNumber", "2008")
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for {
		chunk := make([]byte, 64*1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}
		stream.Send(&messages.AddPhotoRequest{Data: chunk})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.IsOK)

}
