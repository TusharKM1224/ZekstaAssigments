package Repository

import (
	"github.com/TusharKM1224/Types"
	"gorm.io/gorm"
)

type gormDB struct {
	db *gorm.DB
}

type SqlOperations interface{
	CreatNewUser(u*Types.User) error
	CheckEmail(email string) bool

	
	
}

func NewsqlRepository (g *gorm.DB) SqlOperations{
	return &gormDB{db: g}

}
func (con *gormDB) CreatNewUser(u*Types.User) error{
	if err:=con.db.Create(&u).Error;err!=nil{
		return err
	}
	return nil

	
}
func (con *gormDB) CheckEmail(email string) bool{
	var count int64
	if con.db.Model(&Types.User{}).Where("email = ?",email).Count(&count);count>0{
		
		return true
	}
	return false
}