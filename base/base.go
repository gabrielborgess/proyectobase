package base

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Base() { //function principal donde llamamos la funcion de crear tablas etc
	creartablas()
}

func Droptable(nombre string) {
	execdb("DROP TABLE IF EXISTS `" + nombre + "`")
}
func Createtable(nombre string, atributos string) {
	execdb("CREATE TABLE " + nombre + " (" + atributos + ")")
}

func execdb(query string) {
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}
	defer db.Close()
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Query ejecutada correctamente")
}

func creartablas() { //una funcion aparte encargada solo de crear tablas
	//Interfaz()
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}

	fmt.Println("Base de datos conectada exitosamente")

	defer db.Close()
	Droptable("clientes")
	Droptable("vendedores")
	Droptable("pedidos")
	Droptable("detalle_pedidos")
	Droptable("productos")
	Droptable("proveedores")
	Droptable("empleados")
	Droptable("genero")
	Droptable("metodo_pago")
	Createtable("clientes", "id integer, data varchar(32)") // Usa la funcion que cree yo para hacer las query, saldrá mejor y mas facil
	Createtable("vendedores", "id integer, data varchar(32)")
	Createtable("pedidos", "id integer, data varchar(32)")
	Createtable("detalle_pedidos", "id integer, data varchar(32)")
	Createtable("productos", "id integer, data varchar(32)")
	Createtable("proveedores", "id integer, data varchar(32)")
	Createtable("empleados", "id integer, data varchar(32)")
	Createtable("genero", "id integer, data varchar(32)")
	Createtable("metodo_pago", "id integer, data varchar(32)")

	fmt.Println("Tablas Creadas correctamente")

	//execdb("INSERT INTO proveedores VALUES ('0XD')") query mala

}
