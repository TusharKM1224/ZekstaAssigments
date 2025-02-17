package server

import (
	"errors"

	"github.com/TusharKM1224/internals/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBcon struct {
	Db *gorm.DB
}

func NewDBConnection(Dsn string) (*DBcon,error){
	db,err:=gorm.Open(mysql.Open(Dsn),&gorm.Config{})
	if err!=nil{
		return nil,err

	}
	if err:=MigrateModel(db);err!=nil{
		return nil,errors.New("Unable Migrate the structure")
	}
	return &DBcon{Db: db},nil


}

func MigrateModel(db *gorm.DB) error{
	if err:=db.AutoMigrate(&types.DbModel{});err!=nil{
		return err
	}
	return nil
}