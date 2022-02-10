package main

import (
	"context"
	"fmt"
	"hellculator/protobuf"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

type server struct {
	protobuf.UnimplementedCalculatorServiceServer
}

func (s *server) ReqCalc(_ context.Context, in *protobuf.Calc) (*protobuf.Result, error) {
	log.Println("GET ReqCalc with body", in)

	switch in.Operation {
	// `+`
	case protobuf.Calc_ADD:
		return &protobuf.Result{Result: in.A + in.B}, nil

	// `-`
	case protobuf.Calc_SUB:
		return &protobuf.Result{Result: in.A - in.B}, nil

	// `*`
	case protobuf.Calc_MUL:
		return &protobuf.Result{Result: in.A * in.B}, nil

	// `/`
	case protobuf.Calc_DIV:
		if in.B == 0 {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("divide by zero", in.A, in.B))
		}
		return &protobuf.Result{Result: in.A / in.B}, nil

	// if operation not defined.
	default:
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("unknown operation", in.A, in.B))
	}
}

func main() {
	// Listen TCP connection.
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Info("Start listen TCP on port 8080")

	// Init gRPC server.
	grpcServer := grpc.NewServer()

	protobuf.RegisterCalculatorServiceServer(grpcServer, &server{})

	log.Info("gRPC Server Registered")

	// Serve TCP connection.
	log.Info("Start gRPC Serve")
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
