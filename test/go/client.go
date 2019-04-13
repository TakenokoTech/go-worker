package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/TakenokoTech/go-worker/sample"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)
	transform(client)
	stream(client)
}

func transform(client pb.SampleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Transform(ctx, &pb.SampleRequest{Message: "-----"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func stream(client pb.SampleServiceClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Stream(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	defer func() {
		time.Sleep(time.Second)
		log.Printf("close")
		err := stream.CloseSend()
		if err != nil {
			log.Printf("err: %v", err)
		}
	}()

	for i := 1; i <= 10000; i++ {
		log.Printf("index: %v", i)
		err = stream.Send(&pb.SampleRequest{Message: fmt.Sprintf("%v", i)})
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("err: %v", err)
		}

		time.Sleep(time.Millisecond / 120)
	}
}
