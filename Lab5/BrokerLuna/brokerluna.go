package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type BrokerLuna struct {
	pb.UnimplementedBrokerServer
	servers     []string
	vectorClock map[string][]int32
	data        map[string]map[string]int32
}

func newBroker() *BrokerLuna {
	return &BrokerLuna{
		servers:     []string{"localhost:50051", "localhost:50052", "localhost:50053"},
		vectorClock: make(map[string][]int32),
		data:        make(map[string]map[string]int32),
	}
}

func (b *BrokerLuna) GetServerAddress(ctx context.Context, req *pb.AddressRequest) (*pb.ServerAddress, error) {
	if req.Engineer {
		fmt.Printf("Message received from Engineer %d\n", req.Id+1)
	} else {
		fmt.Printf("Message received from Commander\n")
	}
	chosenServer := int32(rand.Intn(3))
	return &pb.ServerAddress{Address: chosenServer}, nil
}

func main() {
	b := newBroker()

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServer(grpcServer, b)

	log.Println("Broker Luna is running on port 50050")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
