package main

import (
	"fmt"
	"sait.mx/primer_programa/model"
	"github.com/gin-gonic/gin"
	"sait.mx/primer_programa/controllers"
)

//con open abres la conexion a bs, regresa dos valores, la conexion y un error, validas que no 
//haya errores

func main() {
	fmt.Println("Pruebas con base de datos")
//todos los metodos van con la primera letra mayuscula!
	model.AbrirDB()

	r := gin.Default()
	r.GET("api/v1/clientes", controllers.ListarClientes) //para obtener la lista de los productos,clientes,objt
	r.GET("api/v1/clientes/:id", controllers.GetCliente) //para obtener el detalle de un cliente, prod, obj
	r.POST("api/v1/clientes", controllers.InsertCliente)
	r.PUT("api/v1/clientes/:id", controllers.UpdateCliente)
	r.DELETE("api/v1/clientes/:id", controllers.BorrarCliente)

	r.Run(":9053")
}


//vprueba para ver si funciono el git 