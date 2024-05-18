package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "La42/Proto"

	"google.golang.org/grpc"
)

type Director struct {
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

func (d *Director) ReportDeath(name string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if status, exists := d.mercenaries[name]; exists {
		status.Alive = false
		fmt.Printf("Mercenary %s has died\n", name)
		// Notify DoshBank (using RabbitMQ, to be implemented)
	}
}

func (d *Director) ReportMove(name string, move *pb.Move) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if status, exists := d.mercenaries[name]; exists && status.Alive {
		// Logic to handle the move and update the game state
		fmt.Printf("Mercenary %s made a move: %v\n", name, move)

		// For example, if move leads to death
		if move.LeadsToDeath {
			d.ReportDeath(name)
		}
	}
}

func (d *Director) CheckStatus() {
	d.mu.Lock()
	defer d.mu.Unlock()

	aliveCount := 0
	for _, status := range d.mercenaries {
		if status.Alive {
			aliveCount++
			fmt.Printf("Mercenary %s is alive\n", status.Name)
		}
	}

	fmt.Printf("%d mercenaries are alive at the end of the floor\n", aliveCount)
}

func (d *Director) AddMercenary(name, address string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.mercenaries[name] = &pb.MercenaryStatus{
		Name:    name,
		Address: address,
		Alive:   true,
		Ready:   false,
	}
}

func (d *Director) RequestAccumulatedAmount() {
	conn, err := grpc.Dial(d.doshBankAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to DoshBank: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewDoshBankClient(conn)
	response, err := client.GetAccumulatedAmount(context.Background(), &pb.Empty{})
	if err != nil {
		log.Printf("Failed to get accumulated amount: %v\n", err)
		return
	}

	fmt.Printf("Accumulated amount: %d\n", response.Amount)
}

func main() {
	director := NewDirector("localhost:50053") // Address of DoshBank service

	// Simulating mercenary addition
	director.AddMercenary("Mercenary1", "localhost:50051")
	director.AddMercenary("Mercenary2", "localhost:50052")

	// Example usage
	director.StartFloor(1)

	// Simulating mercenary moves
	move := &pb.Move{Action: "MoveForward", LeadsToDeath: false}
	director.ReportMove("Mercenary1", move)

	move = &pb.Move{Action: "EnterRoom", LeadsToDeath: true}
	director.ReportMove("Mercenary2", move)

	director.CheckStatus()

	// Request accumulated amount from DoshBank
	director.RequestAccumulatedAmount()
}
