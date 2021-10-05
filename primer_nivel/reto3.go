package main

import (
	"fmt"
	"strconv"
)

/* Suma de enteros
   Instrucciones: otro clásico conocido, donde pedirás al usuario que ingrese 2 números y muestre en
   pantalla el resultado. Si quieres añadir más dificultad puedes utilizar números con punto
   Ejemplo: 2.5633 + 5.6883 = 8.25

*/

func main(){
	var number1 float64
	var number2 float64

	number1 = scanNumber()
	number2 = scanNumber()
	fmt.Println("La suma de números es:")
	fmt.Println(number1 + number2)
}

func scanNumber() float64{
	var lecture string
	importer := false
	var number float64
	for importer == false {
		fmt.Println("inserte un numero:")
		fmt.Scanln(&lecture)
		numero, err := strconv.ParseFloat(lecture, 64)
		if err != nil{
			fmt.Println("Inserte un número válido")
			continue
		}
		number = numero
		importer = true
	}
	return number
}