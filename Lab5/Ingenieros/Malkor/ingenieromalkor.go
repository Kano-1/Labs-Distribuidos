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

func (e *Engineer) sendInformation(server int32, action string, sector string, base string, value string) {
	if _, exists := e.data[sector]; !exists {
		e.data[sector] = make(map[string]int32)
	}

	switch action {
	case "AgregarBase":
		if intValue, err := strconv.Atoi(value); err == nil {
			e.data[sector][base] = int32(intValue)
		}
	case "ActualizarValor":
		if intValue, err := strconv.Atoi(value); err == nil {
			e.data[sector][base] = int32(intValue)
		}
	case "RemombrarBase":
		if oldValue, exists := e.data[sector][base]; exists {
			e.data[sector][value] = oldValue
			delete(e.data[sector], base)
		} else {
			e.data[sector][value] = oldValue
		}

	case "BorrarBase":
		delete(e.data[sector], base)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := e.serverClients[server].WriteInfo(ctx, &pb.ActionRequest{Id: e.id, Action: action, Sector: sector, Base: base, Value: value})
	if err != nil {
		log.Fatalf("Could not send message to Fulcrum Server %d: %v", server+1, err)
	}
	fmt.Printf("Message received from Fulcrum Server %d: [%d, %d, %d]\n", server+1, res.Clock[0], res.Clock[1], res.Clock[2])

	e.vectorClock[sector] = res.Clock
	e.lastServer[sector] = server
}

func (e *Engineer) consoleInterface() {
	var option int32
	var information, action string
	var values []string

	fmt.Println("[Ingeniero Malkor] Comandos posibles:")
	fmt.Println("1. Agregar una base\n2. Actualizar valor de una base\n3. Renombrar una base\n4. Borrar una base")
	fmt.Printf("Seleccione una opción: ")
	fmt.Scan(&option)

	if option == 1 {
		// Agregar una base
		fmt.Printf("Ingrese sector, base y enemigos siguiendo la forma \"<sector>-<base>-<enemigos>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, "-")
		// Caso que no se entregue un número
		if len(values) == 2 {
			values = append(values, "0")
		}
		action = "AgregarBase"
	} else if option == 2 {
		// Actualizar valor de una base
		fmt.Printf("Ingrese sector, base y cantidad actualizada siguiendo la forma \"<sector>-<base>-<cantidad_actualizada>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, "-")
		action = "ActualizarValor"
	} else if option == 3 {
		// Renombrar una base
		fmt.Printf("Ingrese sector, base y nuevo nombre siguiendo la forma \"<sector>-<base>-<nuevo_nombre>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, "-")
		action = "RenombrarBase"
	} else if option == 4 {
		// Borrar una base
		fmt.Printf("Ingrese sector y base siguiendo la forma \"<sector>-<base>\": ")
		fmt.Scan(&information)
		values = strings.Split(information, "-")
		values = append(values, "")
		action = "BorrarBase"
	} else {
		fmt.Println("Opción no válida")
		// Para que no interfiera con lo de abajo
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := e.brokerClient.GetServerAddress(ctx, &pb.AddressRequest{Id: e.id, Engineer: true})
	if err != nil {
		log.Fatalf("Could not send message to Broker Luna: %v", err)
	}
	fmt.Printf("Message received from Broker Luna: %d\n", res.Address)
	e.sendInformation(res.Address, action, values[0], values[1], values[2])
}

func main() {
	e := newEngineer(1)
	for {
		e.consoleInterface()
	}
}
