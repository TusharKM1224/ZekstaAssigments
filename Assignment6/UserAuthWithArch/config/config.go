package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Dbconfig_data struct {
	DB_Host string `validate:"required"`
	DB_Port string `validate:"required"`
	DB_User string `validate:"required"`
	DB_Pass string `validate:"required"`
	DB_Name string `validate:"required"`
	DB_DSN  string `validate:"required"`
}
type Smtpconfig_data struct {
	Org_email     string `validate:"required"`
	Org_App_Pass  string `validate:"required"`
	Org_smtp_host string `validate:"required"`
	Org_smtp_port string `validate:"required"`
}

type Configs_data struct {
	DB_configs   Dbconfig_data
	Smtp_configs Smtpconfig_data
}

func Loadconfig() {
	files := []string{"/home/yuka/ZekstaAssigments/Assignment6/UserAuthWithArch/config/DBconfig.env", "/home/yuka/ZekstaAssigments/Assignment6/UserAuthWithArch/config/OrgConfig.env"}
	for _, file := range files {
		err := godotenv.Load(file)
		if err != nil {
			panic("config file not found!")
		}

	}
}
func Configs() Configs_data {
	Configs_data := Configs_data{DB_configs: Dbconfig_data{
		DB_Host: os.Getenv("DB_HOST"),
		DB_Port: os.Getenv("DB_PORT"),
		DB_User: os.Getenv("DB_USER"),
		DB_Pass: os.Getenv("DB_PASSWORD"),
		DB_Name: os.Getenv("DB_NAME"),
		DB_DSN:  os.Getenv("DB_DSN"),
	},
		Smtp_configs: Smtpconfig_data{
			Org_email:     os.Getenv("ORG_EMAIL"),
			Org_App_Pass:  os.Getenv("ORG_APP_PASS"),
			Org_smtp_host: os.Getenv("ORG_SMTP_HOST"),
			Org_smtp_port: os.Getenv("ORG_SMTP_PORT"),
		}}
	if err := ValidateFields(Configs_data); err != nil {
		panic(err.Error())
	}
	return Configs_data

}

func ValidateFields(C Configs_data) error {

	val := validator.New(validator.WithRequiredStructEnabled())
	if err := val.Struct(C); err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("Missing Field: %s | Error: %s | Value: %v\n", e.Field(), e.Tag(), e.Value()))
		}
		panic("Configuration validation failed!") // Stop execution if config is invalid
	}
	return nil

}
