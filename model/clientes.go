package model

import(
	"fmt"
	_ "github.com/go-sql-driver/mysql" //
	"github.com/jmoiron/sqlx" //sql extended

)

type Cliente struct {
	ID       string `db:"id"         json:"id"` //como estan en la base de datos y como queremos que lo regrese como json
	Nombre   string `db:"nombre"     json:"nombre"` //las de despues de string son etiquetas en las estructuras
	Apellido string `db:"apellido"   json:"apellido"` //alias que indica los cambos de la base de datos
}

var DB *sqlx.DB

//abrir la conexion a la base de datos
func AbrirDB(){
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3307)/pruebasgo")//cambiar root: @ a usuario y clave propio
	if err != nil {
		fmt.Println(err)
		return
	}
	//se asigna la conexion a la variable global
	DB = db
	//ping es para cerciorarme que esta cobectado correctamente
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//recibe un id , regresa un objeto cliente de lo que consulto en la bd y un error
func GetCliente(clientID string) (cliente Cliente, err error) {
	//con get obtienes un solo renglon
	//el & es un apuntador que afecta a la variable
	err = DB.Get(&cliente, `SELECT id, nombre, apellido FROM clientes
		WHERE id=?`, clientID) //en el signo de interrogacion mandamos el parametro
	return
}

//puedes urar nameExec, busca la documentacion de slqx
func Insertar(cliente Cliente) (err error){
	//estoy omitiendo la id que retorne
	_, err = DB.NamedExec(`Insert into clientes (nombre, apellido) values (:nombre, :apellido)`,
	&cliente) //los dos puntos hacen referencia a los campos de la estructura
	return 
}
//delete update y list con bd 
func ListCliente()(arreglo []Cliente, err error){
	err = DB.Select(&arreglo, `Select * from clientes`)
	return
}

func Eliminar(clienteID string) (err error){
	stmt,err := DB.Prepare(`delete from clientes where id =?`)
	stmt.Exec(clienteID)
	return
}

func Update(cliente Cliente) (err error){
	_,err= DB.NamedExec(`update clientes set nombre = :nombre, apellido = :apellido where id = :id`, cliente)
//el funcionamiento del de arriba, solo le esta mandando el objeto pero tiene las etiquetas que estan definidas en la estructura para 
//que automaticamente las cargue de la estructura sin necesidad de hacer nuevas variables, para eso tiene las etiquetas arribaaa
	/*
	nam := cliente.Nombre
	ap := cliente.Apellido
	id:= cliente.ID
	_, err = DB.Query(`update clientes set nombre = ?, apellido = ? where id = ?`, nam, ap, id)*/
	return
}