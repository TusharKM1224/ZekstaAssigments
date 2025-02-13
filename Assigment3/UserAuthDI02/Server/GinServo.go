package Server

import (
	Repository "github.com/TusharKM1224/Repository/MysqlOperations"
	"github.com/TusharKM1224/Routes"
	"github.com/TusharKM1224/Service"
	"github.com/gin-gonic/gin"
)





func InitializeGinServer() (*gin.Engine,Service.Userservice){
	router:=gin.Default()
	dsn:="root:root@tcp(127.0.0.1:3306)/User?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=NewDBConnection(dsn)
	if err!=nil{
		panic(err.Error())
	}
	newrepo:=Repository.NewsqlRepository(db.Db)
	newserviceInstance:=Service.Get_Services(newrepo)
	Routes.UserRoutes(router)

	return router,newserviceInstance
}