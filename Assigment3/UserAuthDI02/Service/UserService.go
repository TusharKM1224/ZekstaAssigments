package Service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"

	"fmt"

	Repository "github.com/TusharKM1224/Repository/MysqlOperations"
	"github.com/TusharKM1224/Types"
)


var DataCache = cache.New(5*time.Minute, 10*time.Minute)

type Repocon struct {
	repo Repository.SqlOperations
}
type Userservice interface {
	PasswordEncryption(pwd string) string
	VerifyUser(U *Types.User) error
	VerifyEmail(Email string, otp string) bool
	generateOtp() string
	ValidateUser(V *Types.ValidateOtp) error
}

func Get_Services(r Repository.SqlOperations) Userservice {
	return &Repocon{repo: r}
}

func (R *Repocon) PasswordEncryption(pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))
}

func (R *Repocon) generateOtp() string {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		panic("Issue with Otp generation")
	}
	return fmt.Sprintf("%06d", (int(b[0])<<16|int(b[1])<<8|int(b[2]))%1000000)

}

func (R *Repocon) VerifyEmail(REmail string, otp string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	senderEmail := "hhornet823@gmail.com"
	senderPassword := "xgce pigo mlyj johw"
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", REmail)
	m.SetHeader("Subject", "Your Otp : ")
	m.SetBody("text/plain", fmt.Sprintf("OTP -> %s", otp))
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		return false
	}
	return true

}

func (R *Repocon) VerifyUser(U *Types.User) error {

	//check in DB
	if isthere := R.repo.CheckEmail(U.Email); isthere {
		return errors.New("Email Already taken")
	}
	Otp := R.generateOtp()
	DataCache.Set(U.Email, map[string]string{
		"otp":       Otp,
		"Name":      U.User_Name,
		"Password":  U.Password,
		"Email":     U.Email,
		"Phone":     U.Phone,
		"Role_name": U.Role_name,
	}, cache.DefaultExpiration)

	if success := R.VerifyEmail(U.Email, Otp); success {
		return nil
	}
	return errors.New("Something went worng!")

}
func (R *Repocon) ValidateUser(V *Types.ValidateOtp) error {
	storedData, found := DataCache.Get(V.Email)
	if !found {
		return errors.New("Otp Expired or invalid")
	}
	Data := storedData.(map[string]string)
	storeOtp := Data["otp"]
	if storeOtp != V.Otp {
		return errors.New("Invalid Otp")
	}
	encryp := R.PasswordEncryption(Data["Password"])
	newuser := &Types.User{User_Name: Data["Name"], Phone: Data["Phone"], Email: Data["Email"], Role_name: Data["Role_name"], Password: encryp}
	if err := R.repo.CreatNewUser(newuser); err != nil {
		return errors.New("Unable to Insert Data")
	}
	return nil

}
