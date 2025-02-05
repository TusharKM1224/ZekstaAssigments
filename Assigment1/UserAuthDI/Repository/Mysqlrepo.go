package Repository

import (
	"context"
	

	Payloadschemas "github.com/TusharKM1224/UserAuth/PayloadSchemas"
	"github.com/TusharKM1224/UserAuth/Types"
	"gorm.io/gorm"
)

type mysqlDBCon struct {
	db *gorm.DB

}

type MysqlOPS interface{
	Create(ctx context.Context,U*Types.UserAuthstruct) error
	UpdateEmailByID(ctx context.Context,U*Payloadschemas.UpdateEmailByIDPayload) error
	GetIdByEmail(ctx context.Context,Email string) bool
	DeleteUser(ctx context.Context,ID string) error
	GetUserbyID(ctx context.Context,ID string) bool
}

func GetmysqlConInstance(Db*gorm.DB) MysqlOPS{
	return &mysqlDBCon{db: Db}
}
func (con*mysqlDBCon) Create(ctx context.Context,U*Types.UserAuthstruct) error{
	tx:=con.db.WithContext(ctx).Begin()
	err:=tx.Create(&U).Error
	if ctx.Err() !=nil{
		tx.Rollback()
		
		return ctx.Err()
	}

	if err!=nil{
		
		tx.Rollback()
		panic(err.Error())
		return err
	}

	tx.Commit()
	return nil
	
}

func (con*mysqlDBCon) UpdateEmailByID(ctx context.Context,U*Payloadschemas.UpdateEmailByIDPayload) error{
	tx:=con.db.WithContext(ctx).Begin()

	err:=tx.Model(&Types.UserAuthstruct{}).Where("id=?",U.ID).Update("email",U.NewEmail).Error
	if err!=nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (con*mysqlDBCon) GetIdByEmail(ctx context.Context,email string) bool{

	var count int64 =0
	tx:=con.db.WithContext(ctx).Begin()
	tx.Model(&Types.UserAuthstruct{}).Where("email=?",email).Count(&count)
	if count>0{
		
		return true
	}

	tx.Commit()
	return false
	
}

func (con*mysqlDBCon) DeleteUser (ctx context.Context,ID string) error{
	tx:=con.db.WithContext(ctx).Begin()
	if err:=tx.Delete(&Types.UserAuthstruct{},ID).Error;err!=nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (con*mysqlDBCon)GetUserbyID(ctx context.Context,ID string) bool {
	
	var count int64 =0
	tx:=con.db.WithContext(ctx).Begin()
	tx.Model(&Types.UserAuthstruct{}).Where("id=?",ID).Count(&count)
	if count>0{
		
		return true
	}

	tx.Commit()
	return false

}