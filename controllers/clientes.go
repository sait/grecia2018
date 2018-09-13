package controllers

import(
	"github.com/gin-gonic/gin"
	"sait.mx/primer_programa/model"
	"io/ioutil"
	"encoding/json"
)

//si regresa status 200 = ok ; 500 = error del servidor
func ListarClientes(c *gin.Context) {
	clientes, err := model.ListCliente()
	if err!=nil {
		c.JSON(500,"Error en el servidor")
		return
	}
	c.JSON(200,clientes)
}

func GetCliente(c *gin.Context){
	id := c.Param("id") // es una variable
	cliente, err := model.GetCliente(id)
	if err!=nil {
		c.JSON(500,"Error en el servidor")
		return
	}
	c.JSON(200, cliente)
}

func InsertCliente(c *gin.Context){
	body, err := ioutil.ReadAll(c.Request.Body) //retorna dos valores, un arreglo de bytes y un error, los bytes es el cuerpo
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente) //interpreta el json y lo transforma en un cliente
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	err = model.Insertar(cliente)
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Insertado correctamente")
}

func UpdateCliente(c *gin.Context){
	id:= c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body) //retorna dos valores, un arreglo de bytes y un error, los bytes es el cuerpo
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente) //interpreta el json y lo transforma en un cliente
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	cliente.ID = id
	err = model.Update(cliente)
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Modificado correctamente")	
}

func BorrarCliente(c *gin.Context){
	id:= c.Param("id")
	err := model.Eliminar(id)
	if err!=nil{
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Eliminado correctamente")		
}
