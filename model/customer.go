package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerID      string `gorm:"type:varchar(255);not null;unique;column:customer_id"`
	CustomerName    string `gorm:"type:varchar(255);not null;column:customer_name"`
	CustomerAddress string `gorm:"type:varchar(255);not null;column:customer_address"`
}
