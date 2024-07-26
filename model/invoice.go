package model

import (
	"time"
)

// Invoice represents the invoice table in the database
type Invoice struct {
	ID         uint          `gorm:"primaryKey;autoIncrement"`
	InvoiceID  string        `gorm:"type:varchar(255);not null;unique"`
	Subject    string        `gorm:"type:varchar(255);not null"`
	CustomerID string        `gorm:"type:varchar(255);not null"`
	Customer   Customer      `gorm:"foreignKey:CustomerID;references:CustomerID"`
	IssueDate  time.Time     `gorm:"type:datetime;not null"`
	DueDate    time.Time     `gorm:"type:datetime;not null"`
	SubTotal   float64       `gorm:"type:decimal(10,2);not null"`
	Tax        float64       `gorm:"type:decimal(10,2);not null"`
	GrandTotal float64       `gorm:"type:decimal(10,2);not null"`
	Status     string        `gorm:"type:varchar(255);not null"`
	Items      []InvoiceItem `gorm:"foreignKey:InvoiceID;references:InvoiceID"`
	CreatedAt  time.Time     `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"type:timestamp;autoUpdateTime"`
}

// InvoiceItem represents the invoice_items table in the database
type InvoiceItem struct {
	ID         uint    `gorm:"primaryKey"`
	InvoiceID  string  `gorm:"not null"`
	ItemID     string  `gorm:"not null"`
	Item       Item    `gorm:"foreignKey:ItemID"`
	Quantity   int     `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
}
