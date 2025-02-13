package main

import (
	"github.com/TusharKM1224/Handlers"
	"github.com/TusharKM1224/Server"
	"github.com/TusharKM1224/Service"
	"github.com/gin-gonic/gin"
)

// USERAUTH02 USING GORM AND GIN WITH DEPENDENCY INJECTION
var newserviceInstance Service.Userservice
func main(){
	var router *gin.Engine
	router, newserviceInstance=Server.InitializeGinServer()
	Handlers.GetServiceInstance(newserviceInstance)
	router.Run(":8080")
	

}