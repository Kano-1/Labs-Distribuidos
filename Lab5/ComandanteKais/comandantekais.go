package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type Commander struct {
	data          map[string]map[string]int32
	lastServer    map[string]int32
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
}

func newCommander() *Commander {
	brokerConn, err := grpc.Dial("localhost:50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to Broker Luna: %v", err)
	}

	var serverClients []pb.ServersClient
	for i, addr := range []string{"localhost:50051", "localhost:50052", "localhost:50053"} {
		serverConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect to Server Fulcrum %d: %v", i, err)
		}
		serverClients = append(serverClients, pb.NewServersClient(serverConn))
	}

	return &Commander{
		data:          make(map[string]map[string]int32),
		lastServer:    make(map[string]int32),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: serverClients,
	}
}

func (c *Commander) requestInformation(server int32, sector string, base string) {
	if _, exists := c.data[sector]; !exists {
		c.data[sector] = make(map[string]int32)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := c.serverClients[server].ReadInfo(ctx, &pb.ReadRequest{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("Could not send message to Server Fulcrum %d: %v", server+1, err)
	}

	c.data[sector][base] = res.Enemies
	c.vectorClock[sector] = res.Clock
	c.lastServer[sector] = server
}

func (c *Commander) consoleInterface() {
	var option int32
	var information string
	var values []string

	fmt.Println("[Comandante Kais] Comandos posibles:")
	fmt.Println("1. Preguntar cantidad de enemigos\n2. Salir")
	fmt.Printf("Seleccione una opción: ")
	fmt.Scan(&option)
	switch option {
	case 1:
		fmt.Printf("Ingrese sector y base siguiendo la forma \"<sector>-<base>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, "-")

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		chosenServer, err := c.brokerClient.GetServerAddress(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("Could not send message to Broker Luna: %v", err)
		}

		c.requestInformation(chosenServer.Address, values[0], values[1])
	}
}

func main() {
	c := newCommander()
	for {
		c.consoleInterface()
	}
}
