package main

import (
	pb "github.com/ant1po1e/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got	request with names : %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
