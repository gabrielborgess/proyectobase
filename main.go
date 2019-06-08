package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:22)/admin_proyecto")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
