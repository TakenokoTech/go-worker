package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/TakenokoTech/go-worker/handlers"
	"github.com/TakenokoTech/go-worker/sample"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start")
	// startRestAPI()
	startGRPC()
}

func startRestAPI() {
	http.HandleFunc("/", handlers.SyncHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func startGRPC() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &handlers.SampleService{})
	s.Serve(l)
}
