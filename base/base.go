package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Base() { //function principal donde llamamos la funcion de crear tablas etc
	go creartablas() // COÑO DE LA MADRE, JODEME QUE ASI DE FACIL ES CREAR UN THREAD ACA !?
	time.Sleep(4 * time.Second)
}

func Droptable(nombre string) {
	execdb("DROP TABLE IF EXISTS `" + nombre + "`")
}
func Createtable(nombre string, atributos string) {
	execdb("CREATE TABLE " + nombre + " (" + atributos + ")")
}

func execdb(query string) { // Usa la funcion que cree yo para hacer las query, saldrá mejor y mas facil
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
	//fmt.Println("Query ejecutada correctamente") la comento por que el output es infumable
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
	go Droptable("clientes")
	go Droptable("vendedores")
	go Droptable("pedidos")
	go Droptable("detalle_pedidos")
	go Droptable("productos")
	go Droptable("proveedores")
	go Droptable("empleados")
	go Droptable("genero")
	go Droptable("metodo_pago")
	time.Sleep(1 * time.Second)
	go Createtable("clientes", "id integer, data varchar(32)")
	go Createtable("vendedores", "id integer, data varchar(32)")
	go Createtable("pedidos", "id integer, data varchar(32)")
	go Createtable("detalle_pedidos", "id integer, data varchar(32)")
	go Createtable("productos", "id integer, data varchar(32)")
	go Createtable("proveedores", "id integer, data varchar(32)")
	go Createtable("empleados", "id integer, data varchar(32)")
	go Createtable("genero", "id integer, data varchar(32)")
	go Createtable("metodo_pago", "id integer, data varchar(32)")
	time.Sleep(2 * time.Second)
	fmt.Println("Tablas Creadas correctamente")

	//execdb("INSERT INTO proveedores VALUES ('0XD')") query mala

}
