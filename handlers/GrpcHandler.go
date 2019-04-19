package handlers

import (
	"context"
	"io"
	"log"

	pb "github.com/TakenokoTech/go-worker/sample"
)

type SampleService struct{}

func (s *SampleService) Transform(context context.Context, req *pb.SampleRequest) (*pb.SampleResponse, error) {
	log.Println("[Transf] Start Receive")
	// log.Println("[Transf] Receive message>> ", req.Message)
	rsp := new(pb.SampleResponse)
	rsp.Message = req.Message
	return rsp, nil
}

func (s *SampleService) Stream(stream pb.SampleService_StreamServer) error {
	for {

		// Recive
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}
		log.Println("[Stream] Start Receive")
		//log.Println("[Stream] Receive message>> ", req.Message)

		// Send
		err = stream.Send(&pb.SampleResponse{Message: req.Message})
		if err != nil {
			log.Printf("err: %v", err)
		}
	}
}
