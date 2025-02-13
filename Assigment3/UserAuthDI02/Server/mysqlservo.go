package Server

import (
	"errors"

	"github.com/TusharKM1224/Types"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dbcon struct{
	Db *gorm.DB
}

func NewDBConnection(dsn string) (*Dbcon,error){
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return nil,errors.New("Unable to establish Connection with Msql Database")

	}
	if err:=migrateDB(db);err!=nil{
		return nil, errors.New("Something went wrong during migration")

	}
	return &Dbcon{Db: db},nil

}

func migrateDB (db*gorm.DB) error{
	if err:=db.AutoMigrate(&Types.User{}); err!=nil{
		return err
	}
	return nil
}