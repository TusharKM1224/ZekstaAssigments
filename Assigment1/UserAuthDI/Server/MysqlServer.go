package Server

import ("gorm.io/gorm"
"gorm.io/driver/mysql"
"github.com/TusharKM1224/UserAuth/Types")

var DB*gorm.DB

func GetDBCon() *gorm.DB{
	return DB
}

func MysqlConnection (){
	// creating Database connection
	dsn:="root:root@tcp(127.0.0.1:3306)/User?charset=utf8mb4&parseTime=True&loc=Local"
	//error var
	var err error
	DB,err=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		panic("Unable establish Database connection!")
		return
	}

	//migrate table struct
	if err=DB.AutoMigrate(&Types.UserAuthstruct{});err!=nil{
		panic("Failed to migrate!")
		return
	}

}