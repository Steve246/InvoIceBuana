package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID      string `gorm:"type:varchar(255);not null;unique"`
	CustomerName    string `gorm:"type:varchar(255);not null"`
	CustomerAddress string `gorm:"type:varchar(255);not null"`
	gorm.Model
}
