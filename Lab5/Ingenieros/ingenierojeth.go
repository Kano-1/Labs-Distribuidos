package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type Engineer struct {
	id            int32
	data          map[string]map[string]int32
	lastServer    map[string]int32
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
}

func newEngineer(id int32) *Engineer {
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

	return &Engineer{
		id:            id,
		data:          make(map[string]map[string]int32),
		lastServer:    make(map[string]int32),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: serverClients,
	}
}

func (e *Engineer) SendMessageToBroker(message string) int32 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Empty{}
	chosenServer, err := e.brokerClient.GetServerAddress(ctx, req)
	if err != nil {
		log.Fatalf("Could not send message to Broker Luna: %v", err)
	}
	return chosenServer.Address
}

func (e *Engineer) sendInformation(server int32, action string, sector string, base string, value string) {
	switch action {
	case "AgregarBase":
		if _, exists := e.data[sector]; !exists {
			e.data[sector] = make(map[string]int32)
			if intValue, err := strconv.Atoi(value); err == nil {
				e.data[sector][base] = int32(intValue)
			}
		}
	case "ActualizarValor":
		if intValue, err := strconv.Atoi(value); err == nil {
			e.data[sector][base] = int32(intValue)
		}
	case "RemombrarBase":
		e.data[sector][value] = e.data[sector][base]
		delete(e.data[sector], base)
	case "BorrarBase":
		delete(e.data[sector], base)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	clock, err := e.serverClients[server].WriteInfo(ctx, &pb.ActionRequest{Action: action, Sector: sector, Base: base, Value: value})
	if err != nil {
		log.Fatalf("Could not send message to Server Fulcrum %d: %v", server+1, err)
	}

	e.vectorClock[sector] = clock.Clock
	e.lastServer[sector] = server
}

func (e *Engineer) consoleInterface() {
	var option int32
	var information, action string
	var values []string

	fmt.Println("Comandos posibles:")
	fmt.Println("1. Agregar una base\n2. Actualizar valor de una base\n3. Renombrar una base\n4. Borrar una base")
	fmt.Printf("Seleccione una opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		// Agregar una base
		fmt.Printf("Ingrese sector, base y enemigos siguiendo la forma \"<sector> <base> <enemigos>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, " ")
		action = "AgregarBase"

	case 2:
		// Actualizar valor de una base
		fmt.Printf("Ingrese sector, base y cantidad actualizada siguiendo la forma \"<sector> <base> <cantidad_actualizada>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, " ")
		action = "ActualizarValor"

	case 3:
		// Renombrar una base
		fmt.Printf("Ingrese sector, base y nuevo nombre siguiendo la forma \"<sector> <base> <nuevo_nombre>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, " ")
		action = "RenombrarBase"

	case 4:
		// Borrar una base
		fmt.Printf("Ingrese sector y base siguiendo la forma \"<sector> <base>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, " ")
		values = append(values, "")
		action = "BorrarBase"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chosenServer, err := e.brokerClient.GetServerAddress(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not send message to Broker Luna: %v", err)
	}
	e.sendInformation(chosenServer.Address, action, values[0], values[1], values[2])
}

func main() {
	e := newEngineer(0)
	for {
		e.consoleInterface()
	}
}
