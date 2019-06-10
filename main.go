package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Interfaz()
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")

	if err != nil {
		fmt.Printf("error al conectar")
		return
	}

	fmt.Println("Base de datos conectada exitosamente")

	defer db.Close()

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

	fmt.Println("Tablas Creadas correctamente")

}
