package main

import (
	"fmt"
)

func main() {
	xd := "<joli>"
	opcion := "<Selecciona una opcion:>"
	fmt.Println("Bienvenido al administrador de videojuegos")
	fmt.Println("Que deseas hacer?.\n")
	fmt.Println("1.Añadir un videojuego")
	fmt.Println("2.Editar videojuego")
	fmt.Println("3.Buscar videojuego")
	fmt.Println("4.Salir\n")
	fmt.Print("Selecciona una opcion:")
	fmt.Scanf("%s", &opcion)
	fmt.Print("Selecciona una opcion:")
	fmt.Scanf("%s", &xd)
}
