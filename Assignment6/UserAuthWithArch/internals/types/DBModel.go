package types

import "github.com/google/uuid"

type DbModel struct {
	Id uuid.UUID `gorm:primarykey`
	Username string `gorm:size:100;"not null" binding:"required"`
	Password string `gorm:size:100;"not null"binding:"required"`
	Email string `gorm:unique;"not null"binding:"required"`
	Phone string `gorm:"not null" binding:"required"`
}


