package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	pb "Lab4/Proto"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

// Conectado a la 3era máquina virtual
const rabbitMQURL = "amqp://guest:guest@dist063.inf.santiago.usm.cl:5672/"

type ServicioMercenariosServer struct {
	pb.UnimplementedServicioMercenariosServer
}

func main() {
	rabbit_conn, rabbit_err := amqp.Dial(rabbitMQURL)

	if rabbit_err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", rabbit_err)
	}

	ch, rabbit_err := rabbit_conn.Channel()
	if rabbit_err != nil {
		fmt.Println("Failed to open a channel:", rabbit_err)
	}

	defer ch.Close()

	grpc_lis, grpc_err := net.Listen("tcp", ":50051")
	if grpc_err != nil {
		fmt.Println("Failed to listen:", grpc_err)
	}
	server := grpc.NewServer()
	servicioMercenariosServer := &ServicioMercenariosServer{}
	pb.RegisterServicioMercenariosServer(server, servicioMercenariosServer)
}

func (s *ServicioMercenariosServer) EnviarDecisiones(ctx context.Context, req *pb.ComunicarDecision) (*pb.RespuestaExito, error) {
}

func logicaPisos(piso int) {
	switch piso {
	case 1:
		// escoger las probabilidades
		probabilidad := rand.Intn(101)
		x := rand.Intn(101)
		y := rand.Intn(101-x) + x + 1
		decision := 0
		switch decision {
		case 0:
			// escogio escopeta
			if probabilidad >= 0 && probabilidad <= x {
				// sobrevive
			} else {
				// muere
			}
		case 1:
			// escogio rifle
			if probabilidad >= x && probabilidad <= y {
				// sobrevive
			} else {
				// muere
			}
		case 2:
			// escogio puños
			if probabilidad >= y && probabilidad <= 100 {
				// sobrevive
			} else {
				// muere
			}
		}
	case 2:
		// escoger a o b
	case 3:
		// escoger los 5 números
	}

}
