package handler

import (
	"context"
	"net/http"
	"time"

	Payloadschemas "github.com/TusharKM1224/UserAuth/PayloadSchemas"
	"github.com/TusharKM1224/UserAuth/Service"
	"github.com/TusharKM1224/UserAuth/Types"
	"github.com/gin-gonic/gin"
)
func getContext(c*gin.Context) (context.Context ,context.CancelFunc){
	ctx,cancel:=context.WithTimeout(c.Request.Context(),5*time.Second)
	return ctx,cancel
}

func UpdateHandler(c*gin.Context)  {
	var newpayload Payloadschemas.UpdateEmailByIDPayload
	ctx,cancel:=getContext(c)
	defer cancel()
	err:=c.BindJSON(&newpayload)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"payload did not match",
		})
	}
	err=Service.UpdateEmailByID(&newpayload,ctx)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Update Successfull",
	})




}
func CreateUserHandler(c*gin.Context){
	var NewUser Types.UserAuthstruct
	err:=c.BindJSON(&NewUser)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid Request format",
		})
		return
	}
	ctx,cancel:=getContext(c)	
	defer cancel()
	err=Service.CreateUser(&NewUser,ctx)
	if err!=nil{
		c.IndentedJSON(http.StatusBadGateway,gin.H{
			"message":err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusAccepted,gin.H{
		"message":"Push Data to Database successfully",
	})

}

func DeleteUserHandler(c*gin.Context){
	ID:=c.Param("id")
	ctx,cancel:=getContext(c)
	defer cancel()
	err:=Service.DeleteUSer(ID,ctx)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusBadRequest,gin.H{
		"message":err.Error(),
	})



}
