package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:22)/admin_proyecto")

	if err != nil {
		fmt.Printf("error al conectar")
	}

	fmt.Println("Base de datos conectada exitosamente")

	defer db.Close()

	dato, err2 := db.Query("select * from tablaprueba")
	if err2 != nil {
		fmt.Println("error prro")
	}

	for dato.Next() {
		var rut string
		err3 := dato.Scan(&rut)

		if err3 != nil {
			fmt.Println("error prro2")
		}
		fmt.Println(rut)

	}

}
