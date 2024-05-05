package main

import (
	"bufio"
	"fmt"
	"os"
)

func crearRegistro(mercenario string, piso int32, decisiones []string) {
	archivo, err := os.Create(fmt.Sprintf("%s_%d.txt", mercenario, piso))
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	for _, decision := range decisiones {
		_, err := archivo.WriteString("* " + decision + "\n")
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}
	}
}

func enviarRegistro(mercenario string, piso int32) {
	archivo, err := os.Open(fmt.Sprintf("%s_%d.txt", mercenario, piso))
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	escaner := bufio.NewScanner(archivo)
	var decisiones string

	for escaner.Scan() {
		decisiones += escaner.Text() + "\n"
	}
	if err := escaner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	// envía de vuelta el string decisiones
}

func main() {
	/*
	   conexión con grpc
	*/
}
