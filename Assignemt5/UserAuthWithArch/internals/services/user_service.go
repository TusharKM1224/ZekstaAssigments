package services

import (
	"errors"
	"fmt"
	"time"

	repository "github.com/TusharKM1224/internals/repository/sql"
	"github.com/TusharKM1224/internals/types"
	"github.com/TusharKM1224/internals/utils"
	"github.com/TusharKM1224/pkg"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)
var Datacache = cache.New(5*time.Minute,10*time.Minute)

type repoInstance struct {
	repo repository.SqlOps
	pkg pkg.Messenger
}



type User_Ops interface{
	Signupuser(U*types.DbModel) error
	Validateuser(V*types.ValidateModel) error
	LoggingIn(L *types.LoginModel) error

}



func Getnewservice(r repository.SqlOps, P pkg.Messenger) User_Ops{
	return &repoInstance{repo: r,pkg: P}
}


func (R *repoInstance) Signupuser(U*types.DbModel) error{
	if isthere:=R.repo.CheckForEmail(U.Email);isthere{
		return errors.New("Emal already Exits")
	}
	Otp:=utils.GenerateOtp()
	Datacache.Set(U.Email,map[string]string{
		"otp": Otp,
		"Username":U.Username,
		"Password":U.Password,
		"Email":U.Email,
		"Phone":U.Phone,
	},cache.DefaultExpiration)
	if success:=R.pkg.Mail_sender("OneTimePassword",fmt.Sprintf("Your OTP : ",Otp),U.Email);!success{
		
		return errors.New("There was a Issue with smtp.")
	}
	return nil
	
	//
}
func (R *repoInstance) Validateuser(V*types.ValidateModel) error{

	storedData,found:=Datacache.Get(V.Email)
	if !found{
		return errors.New("Otp expired")
	}
	Data:=storedData.(map[string]string)
	storedOtp:=Data["otp"]
	if storedOtp!=V.Otp{
		return errors.New("Invalid otp")
	}
	encry:=utils.PasswordEncrpytion(Data["Password"])
	newuser:=&types.DbModel{Username: Data["Username"],Password: encry,Email: Data["Email"],Phone: Data["Phone"],Id: uuid.New()}
	if err:=R.repo.CreateUser(newuser);err!=nil{
		return errors.New("Unable to Insert Data to Database")
	}
	return nil

	
}

func (R*repoInstance) LoggingIn(L*types.LoginModel) error{
	if isthere:=R.repo.CheckForEmail(L.Email);!isthere{
		return errors.New("User not found!")
	}
	if R.repo.Getpassword(L.Email)!=utils.PasswordEncrpytion(L.Pass){
		return errors.New("Invalid Credentials")

	}
	return nil
}