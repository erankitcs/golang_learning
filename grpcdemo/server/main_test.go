package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"testing"

	"github.com/erankitcs/golang_learning/grpcdemo/server/pb/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listner := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	messages.RegisterEmployeeServiceServer(server, &employeeServer{})
	go func() {
		if err := server.Serve(listner); err != nil {
			log.Fatal(err)
		}
	}()
	return func(ctx context.Context, s string) (net.Conn, error) {
		return listner.Dial()
	}
}

func TestEmployeeServer_GetByBadgeNumber(t *testing.T) {
	t.Log("Running GetByBadgeNumber test cases")

	emp_2080 := messages.Employee{
		Id:                  1,
		BadgeNumber:         2080,
		FirstName:           "Grace",
		LastName:            "Decker",
		VacationAccrualRate: 2,
		VacationAccrued:     30,
	}

	tests := []struct {
		name        string
		badgeNumber int32
		res         *messages.EmployeeResponse
		errCode     codes.Code
		errMsg      string
	}{
		{
			"Invalid Request with negative Badge number",
			-1,
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("badge number cant be negative, badge recieved %v", -1),
		},
		{
			"Valid request with badge number",
			2080,
			&messages.EmployeeResponse{Employee: &emp_2080},
			codes.OK,
			"",
		},
		{
			"emp not found",
			1111,
			nil,
			codes.NotFound,
			"employee not found",
		},
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := messages.NewEmployeeServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &messages.GetByBadgeNumberRequest{BadgeNumber: tt.badgeNumber}
			response, err := client.GetByBadgeNumber(ctx, request)
			if response != nil {
				if response.GetEmployee().String() != tt.res.GetEmployee().String() {
					t.Error("response: expected", tt.res.GetEmployee(), "received", response.GetEmployee())
				}
			}
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", tt.errCode.String(), "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}

func TestEmployeeServer_GetAll(t *testing.T) {
	t.Log("Running GetAll test cases")

	tests := []struct {
		name    string
		res     int32
		errCode codes.Code
		errMsg  string
	}{
		{
			"Valid response with all Emp",
			3,
			codes.OK,
			"",
		},
	}
	ctx := context.Background()
	///ctx, canFun := context.WithTimeout(ctx, 10*time.Second)
	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//defer canFun()

	client := messages.NewEmployeeServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allemp := []messages.Employee{}
			request := &messages.GetAllRequest{}
			stream, err := client.GetAll(ctx, request)
			for {
				response, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatal(err)
				}
				t.Log("GetAll stream...")
				allemp = append(allemp, *response.GetEmployee())

			}
			if allemp != nil {
				if len(allemp) != int(tt.res) {
					t.Error("response: expected", tt.res, "emp but received", allemp)
				}
			}
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", tt.errCode.String(), "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
