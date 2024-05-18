package main

import (
	"bufio"
	"context"
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
	CONSOLA  = []string{"DAR"} // En caso que se quiera manejar más de uno por consola
	DIRECTOR = "dist061.inf.santiago.usm.cl:50053"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(DIRECTOR, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	// cambiar con el nombre que quede en el proto
	client := pb.NewServicioMercenariosClient(conn)

	var wg sync.WaitGroup
	for _, mercenario := range NOMBRES {
		wg.Add(1)
		go func(mercenario string) {
			defer wg.Done()
			if mercenarioEnConsola(mercenario) {
				misionesMercenarioConsola(mercenario, client)
			} else {
				misionesMercenarioBot(mercenario, client)
			}
		}(mercenario)
	}
	wg.Wait()
}

func misionesMercenarioConsola(mercenario string, client pb.ServicioMercenariosClient) {
	scanner := bufio.NewScanner(os.Stdin)
	var decision string
	for piso := 1; piso <= 3; piso++ {
		var decisiones []string
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
			// enviar el estado al directo
			// esperar inicio de piso
			switch piso {
			case 1:
				fmt.Println("PISO 1 - ENTRADA AL INFIERNO")
				fmt.Printf("[%s] Ingrese el arma seleccionada: ", mercenario)
				scanner.Scan()
				decision = scanner.Text()
				decisiones = append(decisiones, decision)

			case 2:
				fmt.Println("PISO 2 - TRAMPAS Y TRAICIONES")
				fmt.Printf("[%s] Ingrese el pasillo seleccionado: ", mercenario)
				scanner.Scan()
				decision = scanner.Text()
				decisiones = append(decisiones, decision)

			case 3:
				fmt.Println("PISO 3 - CONFRONTACIÓN FINAL")
				for i := 1; i <= 5; i++ {
					fmt.Printf("[%s] Ingrese un número del 1 al 15 para la ronda %d: ", mercenario, i)
					scanner.Scan()
					decision = scanner.Text()
					decisiones = append(decisiones, decision)
				}
			}
		}

		respuesta, err := client.EnviarDecisiones(context.Background(), &pb.ComunicarDecision{Mercenario: mercenario, Decisiones: decisiones, Piso: piso})
		if err != nil {
			fmt.Println("Failed to send message:", err)
		}
		if !respuesta.Exito {
			// idk el mensaje que debería ser
			fmt.Println(fmt.Sprintf("[%s] Fin del juego, moriste en el piso %d.", mercenario, piso))
			return
		}
	}
	// Si llega acá es porque ganó
	fmt.Println(fmt.Sprintf("[%s] Felicidades, llegaste al final de la misión!", mercenario))
}

func misionesMercenarioBot(mercenario string, client pb.ServicioMercenariosClient) {
	var decision string

	for piso := 1; piso <= 3; piso++ {
		var decisiones []string
		// esperar inicio de piso
		switch piso {
		case 1:
			decision = strconv.Itoa(rand.Intn(3))
			decisiones = append(decisiones, decision)

		case 2:
			decision = strconv.Itoa(rand.Intn(2))
			decisiones = append(decisiones, decision)

		case 3:
			var decisiones []string
			for i := 1; i <= 5; i++ {
				decision += strconv.Itoa(int(rand.Intn(15) + 1))
				decisiones = append(decisiones, decision)
			}
		}

		respuesta, err := client.EnviarDecisiones(context.Background(), &pb.ComunicarDecision{Mercenario: mercenario, Decisiones: decisiones, Piso: piso})
		if err != nil {
			fmt.Println("Failed to send message:", err)
		}
		if !respuesta.Exito {
			// idk el mensaje que debería ser
			log.Println(fmt.Sprintf("[%s] Fin del juego, moriste en el piso %d.", mercenario, piso))
			return
		}
	}
}

// Auxiliar por si se quiere más de un mercenario controlado por consola
func mercenarioEnConsola(mercenario string) bool {
	for _, c := range CONSOLA {
		if mercenario == c {
			return true
		}
	}
	return false
}
