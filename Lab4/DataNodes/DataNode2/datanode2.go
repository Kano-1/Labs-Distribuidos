package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "Lab4/Proto"

	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

type server struct {
	pb.UnimplementedNameNodeServer
}

func (s *server) StoreDecision(ctx context.Context, in *pb.Decision) (*pb.Response, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s_%d.txt", in.Mercenary, in.Floor), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	for _, decision := range in.Decisions {
		if _, err := file.WriteString(fmt.Sprintf("* %s\n", decision)); err != nil {
			return nil, err
		}
	}

	return &pb.Response{Message: "Decision stored successfully"}, nil
}

func (s *server) RetrieveDecisions(ctx context.Context, in *pb.MercenaryFloorRequest) (*pb.DecisionsResponse, error) {
	file, err := os.Open(fmt.Sprintf("%s_%d.txt", in.Mercenary, in.Floor))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decisions := []string{}
	var decision string
	for {
		_, err := fmt.Fscanf(file, "* %s\n", &decision)
		if err != nil {
			break
		}
		decisions = append(decisions, decision)
	}

	return &pb.DecisionsResponse{Decisions: decisions}, nil
}

func main() {
	// Start the DataNode gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameNodeServer(s, &server{})
	log.Printf("DataNode server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
