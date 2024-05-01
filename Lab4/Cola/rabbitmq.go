package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ni idea si esto esta bien, esta copiado del ejemplo y el lab de la sofi xd

func main() {
	// conectado a la máquina virtual 3
	conn, err := amqp.Dial("amqp://guest:guest@dist063.inf.santiago.usm.cl:5672/")

	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
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
	}

	fmt.Println(q)
	defer conn.Close()
}
