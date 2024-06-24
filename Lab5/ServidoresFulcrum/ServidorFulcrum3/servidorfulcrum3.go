package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type FulcrumServer struct {
	pb.UnimplementedServersServer
	id            int32
	log           map[string][]string
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
	mu            sync.Mutex
}

func newServer(id int32) *FulcrumServer {
	brokerConn, err := grpc.Dial("localhost:50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	return &FulcrumServer{
		id:            id,
		log:           make(map[string][]string),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: make([]pb.ServersClient, 0),
	}
}

func (s *FulcrumServer) connectToServers() {
	for i, addr := range []string{"localhost:50051", "localhost:50052"} {
		for {
			serverConn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				log.Printf("Could not connect to Fulcrum Server %d: %v. Retrying...", i+1, err)
				time.Sleep(5 * time.Second)
				continue
			}
			s.serverClients = append(s.serverClients, pb.NewServersClient(serverConn))
			log.Printf("Connected to Fulcrum Server %d", i+1)
			break
		}
	}
}

func (s *FulcrumServer) WriteInfo(ctx context.Context, req *pb.ActionRequest) (*pb.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, action, sector, base, value := req.Id, req.Action, req.Sector, req.Base, req.Value
	fmt.Printf("Message received from Engineer %d: %s %s %s %s\n", id+1, action, sector, base, value)
	s.registerAction(action, sector, base, value, false)
	return &pb.ClockResponse{Clock: s.vectorClock[sector]}, nil
}

func (s *FulcrumServer) ReadInfo(ctx context.Context, req *pb.ReadRequest) (*pb.EnemiesResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sector, base := req.Sector, req.Base
	fmt.Printf("Message received from Commander: %s %s\n", sector, base)
	file, err := os.Open(fmt.Sprintf("%s.txt", sector))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var enemies int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.Contains(l, base) {
			line := strings.Split(l, " ")
			if intValue, err := strconv.Atoi(line[2]); err == nil {
				enemies = int32(intValue)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	file.Close()
	return &pb.EnemiesResponse{Enemies: enemies, Clock: s.vectorClock[sector]}, nil
}

func (s *FulcrumServer) PropagateChanges(ctx context.Context, req *pb.PropagationRequest) (*pb.PropagationResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("Logs received from Fulcrum Server %d\n", req.Id+1)

	// Por cada cambio en el log
	var action, sector, base, value string
	for _, log := range req.Log {
		parts := strings.Split(log, " ")
		if len(parts) == 4 {
			action, sector, base, value = parts[0], parts[1], parts[2], parts[3]
		} else {
			action, sector, base = parts[0], parts[1], parts[2]
			value = ""
		}
		s.registerAction(action, sector, base, value, true)
	}

	for _, clock := range req.Clocks {
		if localClock, exists := s.vectorClock[clock.Sector]; exists {
			for i := range localClock {
				// Guarda el valor mayor
				if localClock[i] < clock.Clock[i] {
					localClock[i] = clock.Clock[i]
				}
			}
		} else {
			s.vectorClock[clock.Sector] = clock.Clock
		}
	}

	return &pb.PropagationResponse{Success: true}, nil
}

func (s *FulcrumServer) registerAction(action string, sector string, base string, value string, propagated bool) {
	if action == "AgregarBase" {
		file, err := os.OpenFile(fmt.Sprintf("%s.txt", sector), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(fmt.Sprintf("%s %s %s\n", sector, base, value)); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		if _, exists := s.vectorClock[sector]; !exists {
			s.vectorClock[sector] = make([]int32, 3)
		}
		if _, exists := s.log[sector]; !exists {
			s.log[sector] = make([]string, 0)
		}
	} else {
		file, err := os.Open(fmt.Sprintf("%s.txt", sector))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		var lines []string
		var oldLine string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.Contains(line, base) {
				lines = append(lines, line)
			} else {
				oldLine = line
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		file.Close()

		switch action {
		case "ActualizarValor":
			lines = append(lines, fmt.Sprintf("%s %s %s", sector, base, value))
		case "RenombrarBase":
			options := strings.Split(oldLine, " ")
			lines = append(lines, fmt.Sprintf("%s %s %s", sector, value, options[2]))
		case "BorrarBase":
			// No hace nada
		}

		err = os.WriteFile(fmt.Sprintf("%s.txt", sector), []byte(strings.Join(lines, "\n")+"\n"), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Si no es un cambio de otro servidor que se propagó
	if !propagated {
		s.vectorClock[sector][s.id] += 1
		s.log[sector] = append(s.log[sector], fmt.Sprintf("%s %s %s %s", action, sector, base, value))
	}
}

func (s *FulcrumServer) startPropagation() {
	for {
		time.Sleep(30 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		logEntries := make([]string, 0)
		vectorClocks := make([]*pb.VectorClocks, 0)

		s.mu.Lock()
		for _, entries := range s.log {
			// Agrega todos los logs del sector
			logEntries = append(logEntries, entries...)
		}
		for sector, clock := range s.vectorClock {
			vectorClocks = append(vectorClocks, &pb.VectorClocks{Sector: sector, Clock: clock})
		}
		s.mu.Unlock()

		// Envío al 1 y 2
		for i, client := range s.serverClients {
			_, err := client.PropagateChanges(ctx, &pb.PropagationRequest{Id: s.id, Log: logEntries, Clocks: vectorClocks})
			if err != nil {
				log.Fatalf("Could not propagate changes to Fulcrum Server %d: %v", i+1, err)
			}
		}
		s.mu.Lock()
		s.log = make(map[string][]string)
		s.mu.Unlock()

		cancel()
	}
}

func main() {
	s := newServer(2)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServersServer(grpcServer, s)

	s.connectToServers()

	log.Printf("Fulcrum Server %d is running on port 50053\n", s.id+1)
	go s.startPropagation()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
