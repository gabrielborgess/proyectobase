package main

import (
	"fmt"
)

func main() {
	var name string
	var xd int
	fmt.Println("Bienvenido al administrador de videojuegos")
	fmt.Printf("Indicame tu nombre:")
	fmt.Scanf("%s", &name)
	fmt.Println("Bienvenido", name)
	fmt.Println("Que deseas hacer?.\n")
	fmt.Println("1.Añadir un videojuego")
	fmt.Println("2.Editar videojuego")
	fmt.Println("3.Buscar videojuego")
	fmt.Println("4.Salir\n")
	fmt.Printf("Selecciona una opcion:.\n")
	fmt.Scanf("%s", &xd)
}
