package main

import (
	"fmt"

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
