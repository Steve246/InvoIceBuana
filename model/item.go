package model

import "time"

// type Item struct {
// 	gorm.Model
// 	Item_ID    string  `gorm:"type:varchar(255);not null;unique;column:item_id"`
// 	Item_Name  string  `gorm:"type:varchar(255);not null;column:item_name"`
// 	Item_Type  string  `gorm:"type:varchar(255);not null;column:item_type"`
// 	Item_Price float64 `gorm:"type:decimal(10,2);not null;column:item_price"`
// }

type Item struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Item_ID    string    `gorm:"type:varchar(255);not null;unique;column:item_id"`
	Item_Name  string    `gorm:"type:varchar(255);not null;column:item_name"`
	Item_Type  string    `gorm:"type:varchar(255);not null;column:item_type"`
	Item_Price float64   `gorm:"type:decimal(10,2);not null;column:item_price"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at"`
}
