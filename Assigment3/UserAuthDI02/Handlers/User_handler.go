package Handlers

import (
	"net/http"

	
	"github.com/TusharKM1224/Service"
	"github.com/TusharKM1224/Types"
	"github.com/gin-gonic/gin"
)
var serviceInstance Service.Userservice

func GetServiceInstance (S Service.Userservice){
	serviceInstance=S
}

func User_signup(c *gin.Context) {
	var new_user Types.User
	if err:=c.BindJSON(&new_user);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid Request Format",
		})
	}
	if err:=serviceInstance.VerifyUser(&new_user);err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Check your Email for Otp",
	})
	
}
func NewUser(c*gin.Context){
	var authdata Types.ValidateOtp
	if err:=c.BindJSON(&authdata);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid Request Format",
		})
		return
	}
	if err:=serviceInstance.ValidateUser(&authdata);err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})

	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"New User created",
	})
}