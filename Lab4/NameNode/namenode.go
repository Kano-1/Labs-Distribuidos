package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func direccionesRegistros(mercenario string, piso int32, datanode string) {
	// Agrega la
	archivo, err := os.OpenFile("direcciones.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(fmt.Sprintf("- %s Piso_%d %s\n", mercenario, piso, datanode))
	if err != nil {
		fmt.Println("Failed to write in file:", err)
		return
	}
}

func obtenerRegistro(mercenario string) {
	/*
		tengo que revisar el archivo direcciones txt
		guardar las líneas que incluyan al mercenario
		pedirle al datanode el registro
		guardar en un string todas las decisiones
		enviar el string al director
	*/
	archivo, err := os.Open("direcciones.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer archivo.Close()

	escaner := bufio.NewScanner(archivo)
	var direccion string
	// string que se va a enviar de vuelta al director con todas las decisiones
	var decisiones string

	// por cada línea
	for escaner.Scan() {
		direccion = escaner.Text()
		// si está el mercenario en esa línea
		if strings.Contains(direccion, mercenario) {
			palabras := strings.Split(direccion, " ")
			// quedaría dividido como ["-", "Piso_n", "datanode"]
			decisiones += palabras[1] + "\n"
			/*
				datanode := palabras[2]
				// ahora le pido al datanode que me de el string y lo sumo

			*/
		}
	}
}
