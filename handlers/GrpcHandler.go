package handlers

import (
	"context"
	"io"
	"log"

	pb "github.com/TakenokoTech/go-worker/sample"
)

type SampleService struct{}

func (s *SampleService) Transform(context context.Context, req *pb.SampleRequest) (*pb.SampleResponse, error) {
	log.Println("call from", req.Message)
	rsp := new(pb.SampleResponse)
	rsp.Message = "Hello " + req.Message + "."
	return rsp, nil
}

func (s *SampleService) Stream(stream pb.SampleService_StreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}
		log.Println("Receive message>> ", req.Message)

		// rsp := new(pb.SampleResponse)
		// rsp.Message = "Hello " + req.Message + "."
		// stream.Send(rsp)
	}
}
