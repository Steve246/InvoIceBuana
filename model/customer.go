package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerID      string `gorm:"type:varchar(255);not null;unique"`
	CustomerName    string `gorm:"type:varchar(255);not null"`
	CustomerAddress string `gorm:"type:varchar(255);not null"`
}
