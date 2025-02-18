package handler

import (
	"net/http"

	"github.com/TusharKM1224/internals/services"
	"github.com/TusharKM1224/internals/types"
	"github.com/gin-gonic/gin"
)

type serviceInstance struct{
	ServInstance services.User_Ops
}

type HandleOps interface{
	CreateNewUser(c*gin.Context)
	ValidateUser(c*gin.Context)
	LogInUser(c*gin.Context)
}

func GetserviceInstance(S services.User_Ops) HandleOps{
	return &serviceInstance{ServInstance: S}
}

func (S *serviceInstance) CreateNewUser(c *gin.Context){
	var newuser types.DbModel
	if err:=c.ShouldBindJSON(&newuser);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid format",
		})
		return
	}
	if err:=S.ServInstance.Signupuser(&newuser);err!=nil{
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

func (S *serviceInstance) ValidateUser(c*gin.Context){
	var newvalidation types.ValidateModel
	
	if err:=c.ShouldBindJSON(&newvalidation);err!=nil{
		
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message": "Invalid format",
		})
		return
	}
	if err:=S.ServInstance.Validateuser(&newvalidation);err!=nil{
		c.IndentedJSON(http.StatusConflict,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusConflict,gin.H{
		"message":"User Added Successfuly",
	})

}

func (S *serviceInstance)  LogInUser(c*gin.Context){
	var cred types.LoginModel
	if err:=c.ShouldBindJSON(&cred);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid format",
		})
		return
	}
	if err:=S.ServInstance.LoggingIn(&cred);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Logging Successfull",
	})

	
}
