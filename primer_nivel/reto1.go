package main

import "fmt"

/* Hola Mundo

   Instrucciones: este es un clásico de clásicos, pero haremos un pequeño cambio.
   En lugar de solo imprimir un mensaje en pantalla, pedirás al usuario que digite un
   nombre y mostrarás en pantalla lo siguiente: Hola, [nombre]
*/

func main (){
	var name string
	fmt.Println("Insert your name")
	fmt.Scanln(&name)
	fmt.Printf("Hola, %s", name)
}