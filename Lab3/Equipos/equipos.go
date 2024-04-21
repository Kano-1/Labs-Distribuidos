package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "Lab3/Proto"

	"google.golang.org/grpc"
)

const numEquipos = 4

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice to hold the channels for communication with each Equipo
	channels := make([]chan struct{}, numEquipos)

	// Create a gRPC connection to the Tierra server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial Tierra server: %v", err)
	}
	defer conn.Close()

	// Create a Tierra client
	client := pb.NewServicioMunicionClient(conn)

	// Start the Equipos routines
	for i := 0; i < numEquipos; i++ {
		channels[i] = make(chan struct{})
		go EquipoRoutine(i+1, client, channels[i])
	}

	// Wait for all Equipos routines to finish
	for _, ch := range channels {
		<-ch
	}

	fmt.Println("Todos los Equipos Conquistaron!.")
}

// EquipoRoutine simulates an Equipo making ammo requests to the Tierra server.
func EquipoRoutine(id int, client pb.ServicioMunicionClient, done chan struct{}) {
	// Wait for 10 seconds before starting requests
	time.Sleep(10 * time.Second)

	for {
		// Generate random number of AT and MP ammo to request
		AT := int32(rand.Intn(11) + 20) // Random number between 20 and 30
		MP := int32(rand.Intn(6) + 10)  // Random number between 10 and 15

		// Send the ammo request to the Tierra server
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		response, err := client.SolicitarM(ctx, &pb.SolicitarMunicion{ID: int32(id), AT: AT, MP: MP})
		if err != nil || !response.Success {
			if err != nil {
				fmt.Printf("\nEquipo %d: \nError: %v \nReintentando en 3 segs...\n", id, err)
			} else {
				fmt.Printf("\nEquipo %d: Solicitando %d AT y %d MP -> Resolucion: -- DENEGADA -- \nReintentando en 3 segs...\n", id, AT, MP)
			}
			time.Sleep(3 * time.Second)
			continue
		}

		// Request approved
		fmt.Printf("\nEquipo %d:Solicitando %d AT y %d MP -> Resolucion: -- APROBADA -- \nConquista Exitosa! Cerrando comunicacion.\n", id, AT, MP)
		break
	}

	// Signal that this Equipo has finished
	done <- struct{}{}
}
