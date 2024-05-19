package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	pb "Lab4/Proto"

	"google.golang.org/grpc"
)

type Mercenary struct {
	Name   string
	Alive  bool
	Ready  bool
	Client pb.DirectorClient
}

func NewMercenary(name string, client pb.DirectorClient) *Mercenary {
	return &Mercenary{
		Name:   name,
		Alive:  true,
		Ready:  false,
		Client: client,
	}
}

func (m *Mercenary) Run(wg *sync.WaitGroup, userControlled bool) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)

	for floor := 1; floor <= 3; floor++ {
		if !m.Alive {
			return
		}

		// Wait for user to indicate ready status
		if userControlled {
			fmt.Printf("Mercenary %s, press Enter when ready for floor %d...", m.Name, floor)
			scanner.Scan()
		}

		// Send ready signal to the Director
		m.sendReadySignal()

		// Wait for the Director to start the floor
		fmt.Printf("Mercenary %s is ready for floor %d\n", m.Name, floor)

		// Perform action based on floor
		var action string
		if userControlled {
			fmt.Printf("Choose action for floor %d: ", floor)
			scanner.Scan()
			action = scanner.Text()
		} else {
			action = m.automatedChoice(floor)
		}

		move := &pb.Move{}
		switch floor {
		case 1:
			move.Action = action // escopeta, rifle, puños electricos
		case 2:
			move.Action = action // hallway A or B
		case 3:
			move.Action = fmt.Sprintf("Numbers: %s", action) // 5 random numbers
		}

		// Send move to the Director
		m.sendMove(move)

		// Check if the mercenary is still alive after the move
		m.checkStatus()
	}
}

func (m *Mercenary) sendReadySignal() {
	_, err := m.Client.ReportMove(context.Background(), &pb.MoveInfo{
		Name: m.Name,
		Move: &pb.Move{Action: "Ready"},
	})
	if err != nil {
		log.Printf("Failed to send ready signal for %s: %v\n", m.Name, err)
	}
}

func (m *Mercenary) sendMove(move *pb.Move) {
	_, err := m.Client.ReportMove(context.Background(), &pb.MoveInfo{
		Name: m.Name,
		Move: move,
	})
	if err != nil {
		log.Printf("Failed to send move for %s: %v\n", m.Name, err)
	}
}

func (m *Mercenary) checkStatus() {
	statusResponse, err := m.Client.CheckStatus(context.Background(), &pb.MercenaryInfo{Name: m.Name})
	if err != nil {
		log.Printf("Failed to check status for %s: %v\n", m.Name, err)
		return
	}

	for _, status := range statusResponse.Mercenaries {
		if status.Name == m.Name {
			m.Alive = status.Alive
			m.Ready = status.Ready
			break
		}
	}
}

func (m *Mercenary) automatedChoice(floor int) string {
	switch floor {
	case 1:
		choices := []string{"escopeta", "rifle", "puños electricos"}
		return choices[rand.Intn(len(choices))]
	case 2:
		choices := []string{"hallway A", "hallway B"}
		return choices[rand.Intn(len(choices))]
	case 3:
		numbers := make([]int, 5)
		for i := 0; i < 5; i++ {
			numbers[i] = rand.Intn(15) + 1
		}
		return fmt.Sprintf("%d %d %d %d %d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	}
	return ""
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure()) // Address of the Director service
	if err != nil {
		log.Fatalf("Failed to connect to Director: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewDirectorClient(conn)

	// Create and run automated mercenaries
	for i := 1; i <= 7; i++ {
		mercenary := NewMercenary(fmt.Sprintf("Mercenary%d", i), client)
		wg.Add(1)
		go mercenary.Run(&wg, false)
	}

	// Create and run user-controlled mercenary
	userMercenary := NewMercenary("UserControlledMercenary", client)
	wg.Add(1)
	go userMercenary.Run(&wg, true)

	wg.Wait()
}
