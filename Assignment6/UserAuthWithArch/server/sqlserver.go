package server

import (
	"errors"
	"fmt"

	"github.com/TusharKM1224/config"
	"github.com/TusharKM1224/internals/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBcon struct {
	Db *gorm.DB
}

func NewDBConnection(config config.Configs_data) (*DBcon, error) {
	fmt.Println(fmt.Sprintf(config.DB_configs.DB_DSN, config.DB_configs.DB_User, config.DB_configs.DB_Pass, config.DB_configs.DB_Port, config.DB_configs.DB_Name))
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(config.DB_configs.DB_DSN, config.DB_configs.DB_User, config.DB_configs.DB_Pass, config.DB_configs.DB_Port, config.DB_configs.DB_Name)), &gorm.Config{})
	if err != nil {
		return nil, err

	}
	if err := MigrateModel(db); err != nil {
		return nil, errors.New("Unable Migrate the structure")
	}
	return &DBcon{Db: db}, nil

}

func MigrateModel(db *gorm.DB) error {
	if err := db.AutoMigrate(&types.DbModel{}); err != nil {
		return err
	}
	return nil
}
