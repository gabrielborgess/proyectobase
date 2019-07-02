package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/cheggaaa/pb.v1"
	_ "gopkg.in/cheggaaa/pb.v1"
	"strings"
	"time"
)

func Base() { //function principal donde llamamos la funcion de crear tablas etc
	count := 52
	bar := pb.StartNew(count)
	bar.ShowCounters = false
	bar.ShowElapsedTime = true
	bar.ShowFinalTime = false
	go creartablas(bar)
	time.Sleep(4 * time.Second)
}

func RemoveSlice(a []string, i int) []string { // Util para borrar un indice exacto en un slice
	a = append(a[:i], a[i+1:]...)
	return a
}

// Util para agregar elementos extras a un slice con un separador
// Tambien sirve por ejemplo cuando tienes un slice, no quieres hacer otro con strings.split y solo quieres agregar cosas
func AppendSliceBySplit(slice []string, element string, separator string) []string {
	element_slice := strings.Split(element, separator)
	for i := 0; i < len(element_slice); i++ {
		slice = append(slice, element_slice[i])
	}
	return slice
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
}

func creartablas(a *pb.ProgressBar) { //una funcion aparte encargada solo de crear tablas
	tablas := [25]string{"clientes_nombre", "clientes_rut", "clientes_direccion", "empleados_nombre", "empleados_sueldo",
		"empleados_rut", "empleados_area", "empleados_direccion", "empleados_cargo", "area", "pedidos_direccion",
		"pedidos_empleado", "pedidos_cliente", "pedidos_detalle", "detalle_pedidos",
		"productos_nombre", "productos_valor", "productos_proveedor",
		"productos_genero", "productos_plataforma", "productos_estrellas",
		"plataforma", "proveedores", "genero", "metodo_pago"}
	atributos := [25]string{"id integer, nombre varchar(150)", "id integer, rut integer, dv integer",
		"id integer, direccion varchar(255), region varchar(150), telefono varchar(12)", "id integer, nombre varchar(255)",
		"id integer, sueldo integer", "id integer, rut integer, dv integer", "id integer, area_id integer",
		"id integer, direccion varchar(255), region varchar(150)", "id integer, cargo varchar(255)", "id integer, tipo varchar(255)",
		"id integer, direccion varchar(255)", "id integer, empleado_id integer", "id integer, cliente_id integer",
		"id integer, valor integer, detalle_id integer", "id integer, producto_id integer, cantidad integer", "id integer, nombre varchar(255)",
		"id integer, valor integer", "id integer, proveedor_id integer", "id integer, genero_id integer",
		"id integer, plataforma_id integer", "id integer, estrellas integer", "id integer, plataforma varchar(25)",
		"id integer, nombre varchar(255), direccion varchar(255), telefono varchar(12)", "id integer, genero varchar(32)", "id integer, nombre varchar(32)"}
	db, err := sql.Open("mysql", "admin_admin:ganzo10.@tcp(158.69.60.190:3306)/admin_proyecto")
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}
	a.Increment()
	defer db.Close()
	for i := 0; i < len(tablas); i++ {
		go Droptable(tablas[i])
		a.Increment()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < len(tablas); i++ {
		go Createtable(tablas[i], atributos[i])
		a.Increment()
	}
	a.Increment()
	a.FinishPrint("Programa Cargado\n")
	time.Sleep(2 * time.Second)

	//execdb("INSERT INTO proveedores VALUES ('0XD')") query mala

}
