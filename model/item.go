package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemId    string  `gorm:"type:varchar(255);not null;unique"`
	ItemName  string  `gorm:"type:varchar(255);not null"`
	ItemType  string  `gorm:"type:varchar(255);not null"`
	ItemPrice float64 `gorm:"type:varchar(255);not null"`
}
