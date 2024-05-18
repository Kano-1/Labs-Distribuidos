package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

// Conectado a la 3era máquina virtual
const rabbitMQURL = "amqp://guest:guest@dist063.inf.santiago.usm.cl:5672/"

func main() {
	rabbit_conn, rabbit_err := amqp.Dial(rabbitMQURL)

	if rabbit_err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", rabbit_err)
	}

	ch, rabbit_err := rabbit_conn.Channel()
	if rabbit_err != nil {
		fmt.Println("Failed to apen a channel:", rabbit_err)
	}

	defer ch.Close()
}

func registrarMuerte(mercenario string, piso int32, monto int32) {
	archivo, err := os.OpenFile("monto_acumulado.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(fmt.Sprintf("- %s Piso_%d %d\n", mercenario, piso, monto))
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func aumentarMonto(mercenario string, piso int32, monto int32) int32 {
	monto += 100000000
	registrarMuerte(mercenario, piso, monto)
	return monto
}
