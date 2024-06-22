package main

import (
	"log"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type Commander struct {
	data          map[string]map[string]int32
	lastServer    map[string][]string
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
}

func newCommander() *Commander {
	brokerConn, err := grpc.Dial("localhost:50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	var serverClients []pb.ServersClient
	for _, addr := range []string{"localhost:50051", "localhost:50052", "localhost:50053"} {
		serverConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect to server: %v", err)
		}
		serverClients = append(serverClients, pb.NewServersClient(serverConn))
	}

	return &Commander{
		data:          make(map[string]map[string]int32),
		lastServer:    make(map[string][]string),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: serverClients,
	}
}

func main() {
	c := newCommander()
	c.data["Sector1"] = make(map[string]int32)
}
