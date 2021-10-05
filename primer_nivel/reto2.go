package main

import "fmt"

/* Hola… nombre y apellido:
   Instrucciones: Ahora que sabemos incluir nombres, podemos agregar más datos.
   Intentemos con un apellido para tener algo así: ``Hola, [nombre] [apellido]```
*/

func main (){
	var first_name, last_name string

	fmt.Println("Introduzca su nombre:")
	fmt.Scanln(&first_name)

	fmt.Println("Introduzca su apellido:")
	fmt.Scanln(&last_name)

	fmt.Printf("Hola, %s %s", first_name, last_name)
}