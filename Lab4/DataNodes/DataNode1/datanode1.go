package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	pb "Lab4/Proto"

	"google.golang.org/grpc"
)

var NAMENODE = "amqp://guest:guest@dist063.inf.santiago.usm.cl:5672/"

func main() {
	conn, err := grpc.Dial(NAMENODE, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	// cambiar con el nombre que quede en el proto
	client := pb.NewServicioDataNodesClient(conn)
}

func crearRegistro(mercenario string, piso int32, decisiones []string) {
	// Creo el archivo para guardar las decisiones
	archivo, err := os.Create(fmt.Sprintf("%s_%d.txt", mercenario, piso))
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer archivo.Close()

	for _, decision := range decisiones {
		_, err := archivo.WriteString("* " + decision + "\n")
		if err != nil {
			fmt.Println("Failed to write in file:", err)
			return
		}
	}
}

func enviarRegistro(mercenario string, piso int32) string {
	// Leo el archivo donde se guardaron las decisiones
	archivo, err := os.Open(fmt.Sprintf("%s_%d.txt", mercenario, piso))
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return ""
	}
	defer archivo.Close()

	escaner := bufio.NewScanner(archivo)
	var decisiones string

	for escaner.Scan() {
		decisiones += escaner.Text() + "\n"
	}
	if err := escaner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return ""
	}
	return decisiones
}
