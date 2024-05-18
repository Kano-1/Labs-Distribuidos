package main

import (
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

func main() {

}
