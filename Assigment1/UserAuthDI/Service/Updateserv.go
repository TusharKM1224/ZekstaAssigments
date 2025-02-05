package Service

import (
	"context"
	"errors"

	Payloadschemas "github.com/TusharKM1224/UserAuth/PayloadSchemas"
	"github.com/TusharKM1224/UserAuth/Repository"
	"github.com/TusharKM1224/UserAuth/Types"

	"github.com/TusharKM1224/UserAuth/Server"
)

func UpdateEmailByID(U*Payloadschemas.UpdateEmailByIDPayload,ctx context.Context) error {
	DB:=Server.GetDBCon()
	
	//check new email already exits
	isthere:=Repository.MysqlOPS.GetIdByEmail(Repository.GetmysqlConInstance(DB),ctx,U.NewEmail)
	if isthere{
		
		return errors.New("Already exits")
	}
	err:=Repository.MysqlOPS.UpdateEmailByID(Repository.GetmysqlConInstance(DB),ctx,U)
	if err!=nil{
		return err
	}
	return nil


	



}

func CreateUser (S*Types.UserAuthstruct,ctx context.Context) error{
	
	//check user is there or not
	isthere:=Repository.MysqlOPS.GetIdByEmail(Repository.GetmysqlConInstance(Server.GetDBCon()),ctx,S.Email)
	if isthere{
		return errors.New("User Already Exists")

	}
	err:=Repository.MysqlOPS.Create(Repository.GetmysqlConInstance(Server.GetDBCon()),ctx,S)
	if err!=nil{
		
		return err
	}
	return nil
	

}

func DeleteUSer(ID string,ctx context.Context) error{


	isthere:=Repository.MysqlOPS.GetUserbyID(Repository.GetmysqlConInstance(Server.GetDBCon()),ctx,ID)
	if !isthere{
		return errors.New("User not found") 
	}

	if err:=Repository.MysqlOPS.DeleteUser(Repository.GetmysqlConInstance(Server.GetDBCon()),ctx,ID);err!=nil{
		return errors.New("User not found")
	}
	return errors.New("User deleted Successfuly!")
	
}