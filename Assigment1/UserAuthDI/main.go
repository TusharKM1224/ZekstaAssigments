package main

import (
	

	handler "github.com/TusharKM1224/UserAuth/handler"
	
	"github.com/TusharKM1224/UserAuth/Server"




	"github.com/gin-gonic/gin"
)


func init(){

	Server.MysqlConnection()
}


func main(){
	router:=gin.Default()
	router.POST("/createUser",handler.CreateUserHandler)
	router.POST("/updateUser",handler.UpdateHandler)
	router.GET("/deleteUser/:id",handler.DeleteUserHandler)
	router.Run(":8080")

}

