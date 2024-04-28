package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	mainServerAddress = "localhost:8080"
	minDelay          = 0
	maxDelay          = 15
)

var captains = []string{"C1", "C2", "C3"}
var planets = []string{"PA", "PB", "PC", "PD", "PE", "PF"}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start sending requests
	for {
		// Choose a random captain and planet
		captain := captains[rand.Intn(len(captains))]
		planet := planets[rand.Intn(len(planets))]

		// Print the discovery
		fmt.Printf("Captain %s has found a treasure on planet %s\n", captain, planet)

		// Send request to main server
		sendRequest(planet, captain)

		// Generate a random delay
		delay := rand.Intn(maxDelay-minDelay+1) + minDelay
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func sendRequest(planet, captain string) {
	conn, err := net.Dial("tcp", mainServerAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()

	// Send message to main server
	message := fmt.Sprintf("%s %s\n", planet, captain)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		return
	}

	fmt.Printf("Request sent by captain %s for planet %s\n\n", captain, planet)
}
