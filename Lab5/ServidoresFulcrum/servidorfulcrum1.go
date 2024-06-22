package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type FulcrumServer struct {
	id            int32
	log           map[string][]string
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
}

func newServer(id int32) *FulcrumServer {
	brokerConn, err := grpc.Dial("localhost:50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	var serverClients []pb.ServersClient
	for _, addr := range []string{"localhost:50052", "localhost:50053"} {
		serverConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect to server: %v", err)
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

func (s *FulcrumServer) registerAction(action string, sector string, base string, value string) {
	if action == "AgregarBase" {
		file, err := os.OpenFile(fmt.Sprintf("%s.txt", sector), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		if _, err := file.WriteString(fmt.Sprintf("%s %s %s\n", sector, base, value)); err != nil {
			return
		}

		if _, exists := s.vectorClock[sector]; !exists {
			s.vectorClock[sector] = make([]int32, 3)
		}
		s.vectorClock[sector][s.id] += int32(1)

		if _, exists := s.log[sector]; !exists {
			s.log[sector] = make([]string, 0)
		}
		s.log[sector] = append(s.log[sector], fmt.Sprintf("%s %s %s %s", action, sector, base, value))
		return
	}

	file, err := os.Open(fmt.Sprintf("%s.txt", sector))
	if err != nil {
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
		return
	}

	if action == "ActualizarValor" {
		lines = append(lines, fmt.Sprintf("%s %s %s\n", sector, base, value))
	} else if action == "RenombrarBase" {
		options := strings.Split(oldLine, " ")
		lines = append(lines, fmt.Sprintf("%s %s %s\n", sector, value, options[2]))
	}

	os.WriteFile(fmt.Sprintf("%s.txt", sector), []byte(strings.Join(lines, "\n")), 0644)

	// Asumo que ya existen
	s.vectorClock[sector][s.id] += 1
	s.log[sector] = append(s.log[sector], fmt.Sprintf("%s %s %s %s", action, sector, base, value))
}

func main() {
	s := newServer(0)

	s.registerAction("AgregarBase", "SectorAlpha", "Campamento1", "5")
	s.registerAction("AgregarBase", "SectorBeta", "Campamento1", "6")
	s.registerAction("AgregarBase", "SectorAlpha", "Campamento2", "10")
	s.registerAction("AgregarBase", "SectorBeta", "Campamento2", "11")
	s.registerAction("AgregarBase", "SectorBeta", "Campamento3", "1")
	s.registerAction("AgregarBase", "SectorBeta", "Campamento4", "7")
	s.registerAction("AgregarBase", "SectorAlpha", "Campamento3", "10")
	s.registerAction("AgregarBase", "SectorAlpha", "Campamento4", "9")
	s.registerAction("AgregarBase", "SectorAlpha", "Campamento5", "9")
	s.registerAction("AgregarBase", "SectorBeta", "Campamento5", "8")

	s.registerAction("ActualizarValor", "SectorAlpha", "Campamento3", "0")
	s.registerAction("ActualizarValor", "SectorBeta", "Campamento2", "5")
	s.registerAction("BorrarBase", "SectorAlpha", "Campamento4", "")
	s.registerAction("RenombrarBase", "SectorBeta", "Campamento1", "Base1")
}
