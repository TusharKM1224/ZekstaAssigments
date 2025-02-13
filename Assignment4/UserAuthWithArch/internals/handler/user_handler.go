package handler

import (
	"net/http"

	"github.com/TusharKM1224/internals/services"
	"github.com/TusharKM1224/internals/types"
	"github.com/gin-gonic/gin"
)

var serviceInstance services.User_Ops
func GetserviceInstace(S services.User_Ops){
	serviceInstance=S
}

func CreateNewUser(c *gin.Context){
	var newuser types.DbModel
	if err:=c.ShouldBindJSON(&newuser);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid format",
		})
		return
	}
	if err:=serviceInstance.Signupuser(&newuser);err!=nil{
		c.IndentedJSON(http.StatusServiceUnavailable,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Check your mail for otp",
	})
	//


}

func ValidateUser(c*gin.Context){
	var newvalidation types.ValidateModel
	
	if err:=c.ShouldBindJSON(&newvalidation);err!=nil{
		
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message": "Invalid format",
		})
		return
	}
	if err:=serviceInstance.Validateuser(&newvalidation);err!=nil{
		c.IndentedJSON(http.StatusConflict,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusConflict,gin.H{
		"message":"User Added Successfuly",
	})

}

func LogInUser(c*gin.Context){
	var cred types.LoginModel
	if err:=c.ShouldBindJSON(&cred);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid format",
		})
		return
	}
	if err:=serviceInstance.LoggingIn(&cred);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Logging Successfull",
	})

	
}
