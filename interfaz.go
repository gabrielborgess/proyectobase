package main

import (
	"fmt"
	"os"
)

//Saludos a github xd
func main() {
	fmt.Println("Presiona 1 para correr el programa")
	fmt.Println("Presiona 2 para salir")
	for {
		menusito()
	}
}

func menusito() {

	var input int
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("Tecla invalida prro")
		return
	}
	switch input {
	case 1:
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

		if opcion == "4" {
			fmt.Println("Gracias por usar the master program xd")
			fmt.Println("chao prro")
			os.Exit(2)
		}
		fmt.Print("Selecciona una opcion:")
		fmt.Scanf("%s", &xd)
	case 2:
		fmt.Println("Gracias por usar the master program xd")
		fmt.Println("chao prro")
		os.Exit(2)
	default:
		fmt.Println("def")
	}

}
