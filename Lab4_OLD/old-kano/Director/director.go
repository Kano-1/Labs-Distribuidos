package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "Lab4/Proto"

	"google.golang.org/grpc"
)

type Director struct {
	pb.UnimplementedDirectorServer
	mercenaries     map[string]*pb.MercenaryStatus
	doshBankAddress string
	mu              sync.Mutex
}

func NewDirector(doshBankAddress string) *Director {
	return &Director{
		mercenaries:     make(map[string]*pb.MercenaryStatus),
		doshBankAddress: doshBankAddress,
	}
}

func (d *Director) StartFloor(floor int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	ready := true
	for _, status := range d.mercenaries {
		if !status.Ready {
			ready = false
			break
		}
	}

	if ready {
		fmt.Printf("Starting floor %d\n", floor)
		// Notify all mercenaries about the floor start
		for name, status := range d.mercenaries {
			if status.Alive {
				// Send floor information to mercenaries
				conn, err := grpc.Dial(status.Address, grpc.WithInsecure())
				if err != nil {
					log.Printf("Failed to connect to mercenary %s: %v\n", name, err)
					continue
				}
				client := pb.NewMercenaryClient(conn)
				_, err = client.StartFloor(context.Background(), &pb.FloorInfo{Floor: int32(floor)})
				if err != nil {
					log.Printf("Failed to start floor for mercenary %s: %v\n", name, err)
				}
				conn.Close()
			}
		}
	} else {
		fmt.Println("Not all mercenaries are ready to start the floor")
	}
}

func (d *Director) ReportDeath(ctx context.Context, req *pb.DeathInfo) (*pb.Empty, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if status, exists := d.mercenaries[req.Name]; exists {
		status.Alive = false
		fmt.Printf("Mercenary %s has died\n", req.Name)
		// Notify DoshBank (using RabbitMQ, to be implemented)
	}
	return &pb.Empty{}, nil
}

func (d *Director) ReportMove(ctx context.Context, req *pb.MoveInfo) (*pb.Empty, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if status, exists := d.mercenaries[req.Name]; exists && status.Alive {
		// Logic to handle the move and update the game state
		fmt.Printf("Mercenary %s made a move: %v\n", req.Name, req.Move)

		// For example, if move leads to death
		if req.Move.LeadsToDeath {
			d.ReportDeath(context.Background(), &pb.DeathInfo{Name: req.Name})
		} else {
			status.Ready = true
		}
	}
	return &pb.Empty{}, nil
}

func (d *Director) CheckStatus(ctx context.Context, req *pb.MercenaryInfo) (*pb.StatusResponse, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	response := &pb.StatusResponse{}
	for _, status := range d.mercenaries {
		response.Mercenaries = append(response.Mercenaries, status)
	}

	return response, nil
}

func (d *Director) AddMercenary(ctx context.Context, req *pb.MercenaryInfo) (*pb.Empty, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.mercenaries[req.Name] = &pb.MercenaryStatus{
		Name:    req.Name,
		Address: req.Address,
		Alive:   true,
		Ready:   false,
	}
	return &pb.Empty{}, nil
}

func (d *Director) RequestAccumulatedAmount(ctx context.Context, req *pb.Empty) (*pb.AccumulatedAmountResponse, error) {
	conn, err := grpc.Dial(d.doshBankAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to DoshBank: %v\n", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewDoshBankClient(conn)
	response, err := client.GetAccumulatedAmount(context.Background(), &pb.Empty{})
	if err != nil {
		log.Printf("Failed to get accumulated amount: %v\n", err)
		return nil, err
	}

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen on port 50050: %v", err)
	}

	grpcServer := grpc.NewServer()
	director := NewDirector("localhost:50053")
	pb.RegisterDirectorServer(grpcServer, director)

	go func() {
		for {
			var command string
			fmt.Println("Enter 'start' to start the next floor:")
			fmt.Scanln(&command)
			if command == "start" {
				director.StartFloor(1) // You may want to track the floor number
			}
		}
	}()

	fmt.Println("Director server is running on port 50050")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
