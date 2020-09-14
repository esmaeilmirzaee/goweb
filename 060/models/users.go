package models

import (
	_ "gorm.io/drivers/postgres"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
