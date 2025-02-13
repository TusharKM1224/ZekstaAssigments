package server

import (
	"os"

	repository "github.com/TusharKM1224/internals/repository/sql"
	"github.com/TusharKM1224/internals/routes"
	"github.com/TusharKM1224/internals/services"
	"github.com/TusharKM1224/pkg"
	"github.com/gin-gonic/gin"
)

func InitializeGin() (*gin.Engine,services.User_Ops){
	router:=gin.Default()
	DbInstance,err:=NewDBConnection(os.Getenv("DB_DSN"))
	if err!=nil{
		panic(err.Error())
	}
	repoInstance:=repository.Getnewrepo(DbInstance.Db)
	newpkgInstance:=pkg.GetpkgIntstance()
	serviceInstance:=services.Getnewservice(repoInstance,newpkgInstance)
	
	routes.User_routes(router)
	
	
	
	
	
	return router,serviceInstance
}