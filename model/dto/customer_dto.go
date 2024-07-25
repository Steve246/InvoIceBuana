package dto

// type GetCustomer struct {
// 	CustomerId      string `json:"customerId"`
// 	CustomerName    string `json: "custName"`
// 	CustomerAddress string `json: "custAddress"`
// }

type CreateCustomer struct {
	CustomerName    string `json:"custName"`
	CustomerAddress string `json:"custAddress"`
}

type DisplayCustomer struct {
	CustomerID      string `gorm:"type:varchar(255);not null;unique"`
	CustomerName    string `gorm:"type:varchar(255);not null"`
	CustomerAddress string `gorm:"type:varchar(255);not null"`
}
