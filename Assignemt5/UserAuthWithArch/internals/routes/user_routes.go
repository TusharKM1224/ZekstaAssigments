package routes

import (
	"github.com/TusharKM1224/internals/handler"
	"github.com/gin-gonic/gin"
)

func User_routes(routes*gin.Engine , H handler.HandleOps){
	routes.POST("/user/adduser",H.CreateNewUser)
	routes.POST("/user/validate",H.ValidateUser)
	routes.POST("/user/LogIn",H.LogInUser)

}