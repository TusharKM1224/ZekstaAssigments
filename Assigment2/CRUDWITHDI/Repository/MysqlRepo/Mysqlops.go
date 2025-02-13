package Repository

import (
	Type "github.com/TusharKM1224/Type/MysqlType"
	"gorm.io/gorm"
)

type ConDB struct {
	Db *gorm.DB
}
type DBOperations interface{
	Create(*Type.TableSchema) error
	
}

func Newrepo (db *gorm.DB) DBOperations{
	return &ConDB{Db:db}
}

func (C*ConDB) Create(U*Type.TableSchema) error{
	if err:=C.Db.Create(&U).Error;err!=nil{
		return err
	}
	return nil

}