package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func execdb(a string) {
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}
	defer db.Close()
	_, err = db.Exec(a)
	if err != nil {
		panic(err)
	}
	fmt.Println("Query ejecutada correctamente")
}

func main() {

	//Interfaz()
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}

	fmt.Println("Base de datos conectada exitosamente")

	defer db.Close()
	execdb("DROP TABLE IF EXISTS `clientes`")
<<<<<<< HEAD
	execdb("CREATE TABLE clientes ( id integer, data varchar(32) )")
=======
	_, err = db.Exec("CREATE TABLE clientes ( id integer, data varchar(32) )")
>>>>>>> 8c05a1217e66ab83ce43fe55802793309b53071d
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `vendedores`")
	execdb("CREATE TABLE vendedores ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `pedidos`")
	execdb("CREATE TABLE pedidos( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `detalle_pedidos`")
	execdb("CREATE TABLE detalle_pedidos ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `productos`")
	execdb("CREATE TABLE productos ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `proveedores`")
	execdb("CREATE TABLE proveedores( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `empleados`")
	execdb("CREATE TABLE empleados ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `genero`")
	execdb("CREATE TABLE genero ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	execdb("DROP TABLE IF EXISTS `metodo_pago`")
	execdb("CREATE TABLE metodo_pago ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	fmt.Println("Tablas Creadas correctamente")

}
