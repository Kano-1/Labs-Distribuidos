package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "Lab4/Proto"

	"google.golang.org/grpc"
)

// depende si dejamos los nombres del pdf o no
var (
	NOMBRES  = []string{"David_Alberts", "Tom_Banner", "Ana_Larive", "Mr_Foster", "Oisten_Jagerhorn", "DJ_Scully", "DAR", "Rae_Higgins"}
	CONSOLA  = "DAR"
	DIRECTOR = "dist061.inf.santiago.usm.cl:50053"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(DIRECTOR, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	// cambiar con el nombre que quede en el proto
	client := pb.NewServicioMercenariosClient(conn)

	var wg sync.WaitGroup
	for _, mercenario := range NOMBRES {
		wg.Add(1)
		go func(mercenario string) {
			defer wg.Done()
			if mercenario == CONSOLA {
				misionesMercenarioConsola(client, mercenario)
			} else {
				misionesMercenarioBot(client, mercenario)
			}
		}(mercenario)
	}
	wg.Wait()
}

func misionesMercenarioConsola(mercenario string, client pb.MercenaryServiceClient) {
	scanner := bufio.NewScanner(os.Stdin)
	var decision string
	for piso := 1; piso <= 3; piso++ {
		// Mostrar el menú de opciones para el jugador
		fmt.Printf("[%s] Opciones antes del piso %d:\n", mercenario, piso)
		fmt.Println("1. Ver monto acumulado")
		fmt.Println("2. Enviar estado al director y comenzar el piso")
		fmt.Print("Seleccione una opción: ")
		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			// pedir ver el monto acumulado

		case "2":
			// enviar el estado al director y esperar respuesta para partir

			switch piso {
			case 1:
				fmt.Println("PISO 1 - ENTRADA AL INFIERNO")
				fmt.Printf("[%s] Ingrese el arma seleccionada: ", mercenario)
				scanner.Scan()
				decision = scanner.Text()

			case 2:
				fmt.Println("PISO 2 - TRAMPAS Y TRAICIONES")
				fmt.Printf("[%s] Ingrese el pasillo seleccionado: ", mercenario)
				scanner.Scan()
				decision = scanner.Text()

			case 3:
				fmt.Println("PISO 3 - CONFRONTACIÓN FINAL")
				var numero string
				for i := 1; i <= 5; i++ {
					fmt.Printf("[%s] Ingrese un número del 1 al 15 para la ronda %d: ", mercenario, i)
					scanner.Scan()
					numero = scanner.Text()
					decision += numero + "-"
				}
				decision = strings.TrimSuffix(decision, "-")
			}
		}

		// enviarDecision debería ser una funcion que le envie la decision al director y retorne si sobrevive o no
		if !enviarDecision(client, mercenario, piso, decision) {
			// idk el mensaje que debería ser
			fmt.Printf("[%s] Fin del juego, moriste en el piso %d.\n", mercenario, piso)
			break
		}
	}
}

func misionesMercenarioBot(mercenario string, client pb.MercenaryServiceClient) {
	var decision string
	for piso := 1; piso <= 3; piso++ {
		switch piso {
		case 1:
			decision = strconv.Itoa(rand.Intn(3))

		case 2:
			decision = strconv.Itoa(rand.Intn(2))

		case 3:
			for i := 1; i <= 5; i++ {
				decision += strconv.Itoa(int(rand.Intn(15)+1)) + "-"
			}
			decision = strings.TrimSuffix(decision, "-")
		}

		// enviarDecision debería ser una funcion que le envie la decision al director y retorne si sobrevive o no
		if !enviarDecision(client, mercenario, piso, decision) {
			// idk el mensaje que debería ser
			log.Printf("[%s] Fin del juego, moriste en el piso %d.\n", mercenario, piso)
			return
		}
	}
}
