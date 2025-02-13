package main

import (
	"fmt"

	Repository "github.com/TusharKM1224/Repository/MysqlRepo"
	"github.com/TusharKM1224/Server"
	Service "github.com/TusharKM1224/Service/mysqlserve"
	Type "github.com/TusharKM1224/Type/MysqlType"
	
	
)



func main() {
	dsn:="root:root@tcp(127.0.0.1:3306)/User?charset=utf8mb4&parseTime=True&loc=Local"

	Dbcon,err:=Server.NewDatabase(dsn)
	if err!=nil{
		fmt.Println(err.Error())
	}
	newuser:=&Type.TableSchema{Name: "Mani",Email: "y6y@gmail.com"}
	newRepo:=Repository.Newrepo(Dbcon.Db)

	newservice:=Service.NewService(newRepo)
	
	if err:=newservice.CreateUser(newuser);err!=nil{
		
		fmt.Println(err)
		
	}
	fmt.Println("Success")

	


}