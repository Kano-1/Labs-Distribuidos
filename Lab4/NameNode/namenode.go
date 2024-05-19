package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "Lab4/Proto" // Replace with the actual path to the generated proto package

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var dataNodeAddresses = []string{
	// "localhost:50052",
	// "localhost:50053",
	// "localhost:50054",
	"dist062.inf.santiago.usm.cl:50054",
	"dist063.inf.santiago.usm.cl:50053",
	"dist064.inf.santiago.usm.cl:50054",
}

type server struct {
	pb.UnimplementedNameNodeServer
	dataNodes []pb.NameNodeClient
}

func (s *server) SendDecision(ctx context.Context, in *pb.Decision) (*pb.Response, error) {
	chosenDataNode := rand.Intn(len(s.dataNodes))
	dataNodeClient := s.dataNodes[chosenDataNode]
	_, err := dataNodeClient.StoreDecision(ctx, in)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile("decisions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprintf("%s Floor_%d %s\n", in.Mercenary, in.Floor, dataNodeAddresses[chosenDataNode])); err != nil {
		return nil, err
	}

	return &pb.Response{Message: "Decision stored successfully"}, nil
}

func (s *server) GetDecisions(ctx context.Context, in *pb.MercenaryRequest) (*pb.DecisionsResponse, error) {
	file, err := os.Open("decisions.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open decisions.txt: %w", err)
	}
	defer file.Close()

	decisions := []string{}
	var mercenary string
	var floor int
	var dataNode string
	matchedDataNodes := map[int]string{}

	for {
		_, err := fmt.Fscanf(file, "%s Floor_%d %s\n", &mercenary, &floor, &dataNode)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read from decisions.txt: %w", err)
		}
		if mercenary == in.Mercenary {
			matchedDataNodes[floor] = dataNode
		}
	}

	for floor, dn := range matchedDataNodes {
		var dataNodeClient pb.NameNodeClient
		for i, address := range dataNodeAddresses {
			if address == dn {
				dataNodeClient = s.dataNodes[i]
				break
			}
		}
		if dataNodeClient == nil {
			return nil, fmt.Errorf("could not find data node client for %s", dn)
		}

		res, err := dataNodeClient.RetrieveDecisions(ctx, &pb.MercenaryFloorRequest{Mercenary: in.Mercenary, Floor: int32(floor)})
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve decisions: %w", err)
		}
		decisions = append(decisions, res.Decisions...)
	}

	return &pb.DecisionsResponse{Decisions: decisions}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dataNodes := make([]pb.NameNodeClient, len(dataNodeAddresses))

	// Connect to DataNodes
	for i, address := range dataNodeAddresses {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Can't connect to DataNode at %s: %v", address, err)
		}
		defer conn.Close()
		client := pb.NewNameNodeClient(conn)
		dataNodes[i] = client
	}

	// Start the NameNode gRPC server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameNodeServer(s, &server{
		dataNodes: dataNodes,
	})
	log.Printf("NameNode server listening at %v", listener.Addr())
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Simulate test cases
	time.Sleep(2 * time.Second) // Wait for the server to start

	conn, err := grpc.Dial("dist064.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewNameNodeClient(conn)

	// Test: SendDecision
	_, err = client.SendDecision(context.Background(), &pb.Decision{Mercenary: "Merc1", Floor: 1, Decisions: []string{"Puños electricos"}})
	if err != nil {
		log.Fatalf("could not send decision: %v", err)
	}
	_, err = client.SendDecision(context.Background(), &pb.Decision{Mercenary: "Merc1", Floor: 2, Decisions: []string{"Pasillo A"}})
	if err != nil {
		log.Fatalf("could not send decision: %v", err)
	}
	_, err = client.SendDecision(context.Background(), &pb.Decision{Mercenary: "Merc1", Floor: 3, Decisions: []string{"2", "4", "15", "7", "10"}})
	if err != nil {
		log.Fatalf("could not send decision: %v", err)
	}

	// Wait for some time to ensure the decision is processed
	time.Sleep(5 * time.Second)

	// Test: GetDecisions
	resp, err := client.GetDecisions(context.Background(), &pb.MercenaryRequest{Mercenary: "Merc1"})
	if err != nil {
		log.Fatalf("could not get decisions: %v", err)
	}
	fmt.Printf("Decisions: %v\n", resp.Decisions)
}
