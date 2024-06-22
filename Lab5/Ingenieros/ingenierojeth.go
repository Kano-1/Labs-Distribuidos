package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "Lab5/Proto"

	"google.golang.org/grpc"
)

type Engineer struct {
	id            int32
	data          map[string]map[string]int32
	lastServer    map[string][]string
	vectorClock   map[string][]int32
	brokerClient  pb.BrokerClient
	serverClients []pb.ServersClient
}

func newEngineer(id int32) *Engineer {
	brokerConn, err := grpc.Dial("localhost:50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	var serverClients []pb.ServersClient
	for _, addr := range []string{"localhost:50051", "localhost:50052", "localhost:50053"} {
		serverConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect to server: %v", err)
		}
		serverClients = append(serverClients, pb.NewServersClient(serverConn))
	}

	return &Engineer{
		id:            id,
		data:          make(map[string]map[string]int32),
		lastServer:    make(map[string][]string),
		vectorClock:   make(map[string][]int32),
		brokerClient:  pb.NewBrokerClient(brokerConn),
		serverClients: serverClients,
	}
}

func (e *Engineer) sendInformation(action string, sector string, base string, value string) {
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
	// Envía el mensaje
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
	// llamo a la función
	e.sendInformation(action, values[0], values[1], values[2])

}

func main() {
	e := newEngineer(0)
	for {
		e.consoleInterface()
	}
	// e.data["Sector1"] = make(map[string]int32)
	// e.data["Sector1"]["Base1"] = 10
	// e.sendInformation("RenombrarBase", "Sector1", "Base1", "BaseAlpha")
}
