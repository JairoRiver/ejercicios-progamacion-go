package main

import (
	"fmt"
	"strconv"
)

/*
Suma y Multiplición
El usuario ingresará 3 números, sumarás los 2 primeros y el resultado será
multiplicado por el tercero.
*/

func main(){
	number1 := scanNumber()
	number2 := scanNumber()
	number3 := scanNumber()

	response := (number1 + number2) * number3
	fmt.Printf("La respuesta es %v", response)
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