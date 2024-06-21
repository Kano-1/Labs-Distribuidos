package main

import (
	"fmt"
)

type BrokerLuna struct {
	servers     []string
	vectorClock map[string][]int32
	data        map[string]map[string]int32 // No estoy 100% segura si el broker también guarda información
}

func newBroker() *BrokerLuna {
	return &BrokerLuna{
		servers:     []string{"localhost:50051", "localhost:50052", "localhost:50053"},
		vectorClock: make(map[string][]int32),
		data:        make(map[string]map[string]int32),
	}
}

func main() {
	b := newBroker()
	fmt.Println("Hello World!")
	for _, entry := range b.servers {
		fmt.Println(entry)
	}
}
