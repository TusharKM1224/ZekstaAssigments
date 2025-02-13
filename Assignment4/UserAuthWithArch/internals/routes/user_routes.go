package routes

import (
	"github.com/TusharKM1224/internals/handler"
	"github.com/gin-gonic/gin"
)

func User_routes(routes*gin.Engine){
	routes.POST("/user/adduser",handler.CreateNewUser)
	routes.POST("/user/validate",handler.ValidateUser)
	routes.POST("/user/LogIn",handler.LogInUser)

}