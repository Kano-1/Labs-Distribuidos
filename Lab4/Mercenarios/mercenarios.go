package main

import (
	"fmt"
	"math/rand"
)

/*
antes de cada piso: informar al director de su estado -> sólo el que se controla
tomar decisiones para los 3 pisos
cuando mueren términan su código
pedir ver el monto -> sólo el que se controla
prints están por mientras xd
*/

func primerPiso() {
	// seleccionar arma A, B ó C
	seleccion := int32(rand.Intn(3))
	fmt.Println("Se seleccionó el arma:", seleccion)
	/*
		mandarle al director el arma seleccionada
		el director calcula las probabilidades de las 3 armas
		v1:
			el director devuelve la probabilidad del arma
			el mercenario calcula si sobrevive
			le informa al director si sobrevivió
		v2:
			el director calcula la probabilidad si sobrevive
			le informa al mercenario si sobrevivió
	*/
}

func segundoPiso() {
	// seleccionar camino A ó B
	seleccion := int32(rand.Intn(2))
	fmt.Println("Se seleccionó el pasillo:", seleccion)
	/*
		mandarle al director el pasillo seleccionado
		el director selecciona el pasillo correcto
		v1:
			el director le informa al mercenario si sobrevivió
		v2:
			el director devuelve el pasillo seleccionado
			el mercenario revisa si sobrevive
			le informa al director si sobrevivió
	*/
}

func tercerPiso() {
	numeros := make([]int32, 0, 5)
	for i := 0; i < 5; i++ {
		numeros = append(numeros, int32(rand.Intn(15)+1))
		fmt.Println("Se seleccionó el número", numeros[i], "en la ronda", i)
	}
	/*

	 */
}
