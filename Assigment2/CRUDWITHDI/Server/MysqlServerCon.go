package Server

import (
	"errors"

	Type "github.com/TusharKM1224/Type/MysqlType"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)


type DBCon struct{
	Db *gorm.DB
}


func NewDatabase(dsn string) (*DBCon,error){
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		
		return nil,errors.New("Could not establish database connection")
	}
	if err:=migrate(db);err!=nil{
		return nil,errors.New("Uanble to migrate")
	}
	return &DBCon{Db: db} ,nil
}

func migrate(db *gorm.DB) error{
	if err:=db.AutoMigrate(&Type.TableSchema{}); err!=nil{
		return err
	}
	return nil
}

