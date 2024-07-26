package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_ID    string  `gorm:"type:varchar(255);not null;unique;column:item_id"`
	Item_Name  string  `gorm:"type:varchar(255);not null;column:item_name"`
	Item_Type  string  `gorm:"type:varchar(255);not null;column:item_type"`
	Item_Price float64 `gorm:"type:decimal(10,2);not null;column:item_price"`
}
