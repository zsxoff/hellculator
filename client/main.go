package main

import (
	"context"
	"fmt"
	"hellculator/protobuf"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := protobuf.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	result, err := client.ReqCalc(ctx, &protobuf.Calc{A: 2.0, B: 1.0, Operation: protobuf.Calc_add})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your result is:", result.Result)
}
