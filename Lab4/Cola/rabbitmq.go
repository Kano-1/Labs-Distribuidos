package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Conectado a la 3era máquina virtual
const rabbitMQURL = "amqp://guest:guest@dist063.inf.santiago.usm.cl:5672/"

func main() {
	conn, err := amqp.Dial(rabbitMQURL)

	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
		panic(err)
	}

	fmt.Println(q)
	defer conn.Close()
}
