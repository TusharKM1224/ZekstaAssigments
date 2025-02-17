package repository

import (
	"github.com/TusharKM1224/internals/types"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

type SqlOps interface{
	CreateUser(U*types.DbModel) error
	CheckForEmail(Email string) bool
	Getpassword(Email string) string


}

func Getnewrepo(Db*gorm.DB) SqlOps{
	return &DbInstance{Db: Db}
}

func (con*DbInstance) CreateUser(U*types.DbModel) error{
	//signup
	if err:=con.Db.Create(&U).Error;err!=nil{
		return err
	}
	return nil

}
func (con*DbInstance) CheckForEmail(Email string) bool{
	//check
	var count int64
	if con.Db.Model(&types.DbModel{}).Where("email = ?",Email).Count(&count);count>0{
		
		return true
	}
	return false
}
func (con*DbInstance) Getpassword(Email string) string{
	
	var user types.DbModel
	con.Db.Where("email = ?", Email).First(&user) // SELECT * FROM users WHERE email = ? LIMIT 1;
	return user.Password

//
	
}

