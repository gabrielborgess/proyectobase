package main
//importamos el os
import (
	"fmt"
	"os"
)

//Saludos a github xd

//funcion para iniciar programa
func main() {
	fmt.Println("Presiona 1 para correr el programa")
	fmt.Println("Presiona 2 para salir")
	for {
		menusito()
	}
}

func Consulta()string{//Funcion que sirve para ingresar en una variable lo que el usuario quiere buscar asi lo pasamos a una consulta sql
	var atributo string
	atributo="Aqui va lo que la persona quiere buscar en sql"
	fmt.Print("Ingrese lo que desea buscar:")
	fmt.Scanf("%s",&atributo)
	return atributo


}
//en menucito corremos lo principal
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
		fmt.Println("3.Editar videojuego")
		fmt.Println("4.Buscar videojuego")
		fmt.Println("5.Añadir venta")
		fmt.Println("6.Salir\n")
		fmt.Print("Selecciona una opcion:")
		fmt.Scanf("%s", &opcion)

		if opcion=="4"{
			juego:=Consulta()
			fmt.Println("El juego que quiere buscar es:",juego)//Esta linea la puse para probar la funcion despues se elimina
		}

		if opcion == "6" {//aqui aparece algunos if
			fmt.Println("Gracias por usar the master program xd")
			fmt.Println("chao prro")
			os.Exit(2)
		}
		fmt.Print("Selecciona una opcion:")
		fmt.Scanf("%s", &xd)
	case 2://en el menu principial 1 si presionamos el 2 (case 2) cerramos el programa y lanzamos un mensaje
		fmt.Println("Gracias por usar the master program xd")
		fmt.Println("chao prro")
		os.Exit(2)
	default:
		fmt.Println("def")
	}

}
