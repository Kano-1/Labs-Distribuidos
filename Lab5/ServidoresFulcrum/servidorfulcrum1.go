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

	var serverClients []pb.ServersClient
	for _, addr := range []string{"localhost:50052", "localhost:50053"} {
		serverConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect to server: %v", err)
		}
		serverClients = append(serverClients, pb.NewServersClient(serverConn))
	}

	return &FulcrumServer{
		id:            id,
		log:           make(map[string][]string),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: serverClients,
	}
}

func (s *FulcrumServer) WriteAction(ctx context.Context, req *pb.ActionRequest) (*pb.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	action, sector, base, value := req.Action, req.Sector, req.Base, req.Value
	fmt.Printf("Se recibió el mensaje: %s %s %s %s\n", action, sector, base, value)
	s.registerAction(action, sector, base, value)
	return &pb.ClockResponse{Clock: s.vectorClock[sector]}, nil
}

func (s *FulcrumServer) ReadInfo(ctx context.Context, req *pb.ReadRequest) (*pb.EnemiesResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sector, base := req.Sector, req.Base
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

func (s *FulcrumServer) registerAction(action string, sector string, base string, value string) {
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
	s.vectorClock[sector][s.id] += 1
	s.log[sector] = append(s.log[sector], fmt.Sprintf("%s %s %s %s", action, sector, base, value))
}

func main() {
	s := newServer(0)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServersServer(grpcServer, s)

	log.Printf("Fulcrum Server %d is running on port 50051\n", s.id+1)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	for {
		// Propaga los cambios cada 30 segundos
		time.Sleep(30 * time.Second)
	}
}
