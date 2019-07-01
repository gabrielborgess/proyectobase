package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/cheggaaa/pb.v1"
	_ "gopkg.in/cheggaaa/pb.v1"
	"time"
)

func Base() { //function principal donde llamamos la funcion de crear tablas etc
	go creartablas() // COÑO DE LA MADRE, JODEME QUE ASI DE FACIL ES CREAR UN THREAD ACA !?
	go Barra()
	time.Sleep(4 * time.Second)

}

func Barra() {
	count := 10
	bar := pb.StartNew(count)
	bar.Increment()
	bar.FinishPrint("The End!")
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
	porcentaje := 0
	tablas := [9]string{"clientes", "vendedores", "pedidos", "detalle_pedidos",
		"productos", "proveedores", "empleados", "genero", "metodo_pago"}
	atributos := [9]string{"id integer, data varchar(32)", "id integer, data varchar(32)",
		"id integer, data varchar(32)", "id integer, data varchar(32)", "id integer, data varchar(32)",
		"id integer, data varchar(32)", "id integer, data varchar(32)", "id integer, data varchar(32)",
		"id integer, data varchar(32)"}
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}

	fmt.Println(porcentaje, "%")

	defer db.Close()
	for i := 0; i < len(tablas); i++ {
		go Droptable(tablas[i])
		porcentaje += 5
		fmt.Println(porcentaje, "%")
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < len(tablas); i++ {
		go Createtable(tablas[i], atributos[i])
		porcentaje += 5
		fmt.Println(porcentaje, "%")
	}
	porcentaje += 10
	fmt.Println(porcentaje, "%")
	time.Sleep(2 * time.Second)

	//execdb("INSERT INTO proveedores VALUES ('0XD')") query mala

}
