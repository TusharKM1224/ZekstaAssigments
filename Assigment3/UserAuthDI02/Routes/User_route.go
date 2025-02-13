package Routes

import (
	"github.com/TusharKM1224/Handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine){
	router.POST("/user/signUp",Handlers.User_signup)
	router.POST("/user/validation",Handlers.NewUser)
}