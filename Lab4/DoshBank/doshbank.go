package main

import (
	"fmt"
	"os"
)

func registrarMuerte(mercenario string, piso int32, monto int32) {
	archivo, err := os.OpenFile("registro.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(fmt.Sprintf("%s Piso_%d %d\n", mercenario, piso, monto))
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func main() {

}
