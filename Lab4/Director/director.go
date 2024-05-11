package main

import (
	"context"
	"fmt"
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
