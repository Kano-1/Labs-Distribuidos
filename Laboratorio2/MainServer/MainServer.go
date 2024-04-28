package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

var planets = map[string]int{
	"PA": 0,
	"PB": 0,
	"PC": 0,
	"PD": 0,
	"PE": 0,
	"PF": 0,
}

const port = ":8080"

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize treasures on planets
	initializeTreasures()
	fmt.Println("Initial status of assignments:")
	for name, booty := range planets {
		fmt.Printf("%s: %d  ", name, booty)
	}
	fmt.Printf("\n\n")

	// Start server
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()
	fmt.Printf("Server listening on port %s\n\n", port)

	// Handle incoming messages
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleRequest(conn)
	}
}

func initializeTreasures() {
	for planet := range planets {
		planets[planet] = rand.Intn(11) // Random number between 0 and 10
	}

}

func handleRequest(conn net.Conn) {
	// Close connection when this function ends
	defer conn.Close()

	// Read incoming message
	message := make([]byte, 1024)
	n, err := conn.Read(message)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Process message
	messageStr := strings.TrimSpace(string(message[:n]))
	parts := strings.Split(messageStr, " ")
	if len(parts) != 2 {
		fmt.Println("Invalid message format")
		return
	}
	planet := parts[0]
	captain := parts[1]

	// Print received request
	fmt.Printf("Request received from planet %s by captain %s\n", planet, captain)

	// Assign treasure to planet with lowest amount
	minTreasure := 50
	var minPlanet string
	for p, t := range planets {
		if t < minTreasure {
			minTreasure = t
			minPlanet = p
		} else if t == minTreasure && p < minPlanet {
			minPlanet = p
		}
	}

	// Update planet's treasure
	planets[minPlanet]++
	currentAmount := planets[minPlanet]

	// Print assigned treasure
	fmt.Printf("Treasure assigned to planet %s. Current amount: %d\n", minPlanet, currentAmount)
	fmt.Println("Current status of assignments:")
	for name, booty := range planets {
		fmt.Printf("%s: %d  ", name, booty)
	}
	fmt.Printf("\n\n")
}
