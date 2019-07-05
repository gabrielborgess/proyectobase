package main

//importamos el os
import (
	"./base"
	"fmt"
	"os"
	"time"
)

//Saludos a github xd

//funcion para iniciar programa
// No sé que clase de atrocidad hiciste men xdd
func main() {
	var setup string
	fmt.Println("Bienvenido, Desea Instalar la base de datos? S/n")
	fmt.Scanf("%s", &setup)
	if setup == "s" {
		base.Base()
	}
	for {
		menusito()
	}
}

func Consulta() string { //Funcion que sirve para ingresar en una variable lo que el usuario quiere buscar asi lo pasamos a una consulta sql
	var atributo string
	atributo = "Aqui va lo que la persona quiere buscar en sql"
	fmt.Print("Ingrese lo que desea buscar:")
	fmt.Scanf("%s", &atributo)

	return atributo

}

func helpmenu() {
	fmt.Println("Bienvenido al administrador de videojuegos")
	fmt.Println("Que deseas hacer?.\n")
	fmt.Println("1.Añadir un videojuego")
	fmt.Println("2.Editar videojuego")
	fmt.Println("3.Buscar videojuego")
	fmt.Println("4.Añadir venta")
	fmt.Println("5.Buscar venta")
	fmt.Println("6.Editar venta")
	fmt.Println("7.Salir\n")
	fmt.Printf(":")
}

//en menucito corremos lo principal
func menusito() {
	db, _ := base.ObtenerBaseDeDatos()
	defer db.Close()
	opcion := "<Selecciona una opcion:>"
	go helpmenu()
	fmt.Scanf("%s", &opcion)
	switch opcion {
	case "1":
		var (
			nombre string
		)

		fmt.Printf("Ingrese nombre del juego:")
		fmt.Scanf("%s", &nombre)
		go base.Insertar_sql("productos_nombre", "id,nombre", "1,"+nombre, db, false)

	case "2":
		fmt.Println("elegista la Opcion 2")

	case "3":
		juego := Consulta()
		fmt.Println("El juego que quiere buscar es:", juego) //Esta linea la puse para probar la funcion despues se elimina
	case "4":
		fmt.Print("elegiste la opcion 4")
	case "5":
		fmt.Print("elegiste la opcion 5")
	case "6":
		fmt.Print("elegiste la opcion 6")
	case "7":
		fmt.Println("Gracias por usar the master program xd")
		fmt.Println("chao prro")
		os.Exit(2)

	default:
		fmt.Println("Opcion Invalida")
	}
	time.Sleep(1 * time.Second)
}
