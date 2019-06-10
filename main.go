package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Datos struct {
	Rut      string
	Apellido string
}

func main() {
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")

	if err != nil {
		fmt.Printf("error al conectar")
		return
	}

	fmt.Println("Base de datos conectada exitosamente")

	defer db.Close()

	fmt.Println("Seleccion de tabla:")

	dato, err2 := db.Query("select * from tablaprueba")
	if err2 != nil {
		fmt.Println("error prro")
	}

	for dato.Next() {

		d := Datos{}
		err3 := dato.Scan(&d.Rut)

		if err3 != nil {
			fmt.Println("error prro2")
		}
		fmt.Println(d.Rut)

	}

	_, err = db.Exec("CREATE TABLE clientes ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE vendedores ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE pedidos( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE detalle_pedidos ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE productos ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE proveedores( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE empleados ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE genero ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE metodo_pago ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
}
