package main

import (
	"fmt"
	"strconv"
)

type Engineer struct {
	id          int
	servers     []string
	data        map[string]map[string]int32
	lastServer  map[string][]string
	vectorClock map[string][]int32
}

func newEngineer() *Engineer {
	return &Engineer{
		id:          1,
		servers:     []string{"localhost:50051", "localhost:50052", "localhost:50053"},
		data:        make(map[string]map[string]int32),
		lastServer:  make(map[string][]string),
		vectorClock: make(map[string][]int32),
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

func consoleInterface() {
	var option int32
	fmt.Println("Comandos posibles:")
	fmt.Println("1. Agregar una base")
	fmt.Println("2. Actualizar valor de una base")
	fmt.Println("3. Renombrar una base")
	fmt.Println("4. Borrar una base")
	fmt.Printf("Seleccione una opción: ")
	fmt.Scan(&option)
	switch option {
	case 1:
		// agregar una base
	case 2:
		// actaulizar valor de una base
	case 3:
		// renombrar una base
	case 4:
		// borrar una base
	}
}

func main() {
	e := newEngineer()
	consoleInterface()
	e.data["Sector1"] = make(map[string]int32)
	e.data["Sector1"]["Base1"] = 10
	e.sendInformation("RenombrarBase", "Sector1", "Base1", "BaseAlpha")
}
