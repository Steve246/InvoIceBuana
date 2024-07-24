package dto

// type GetCustomer struct {
// 	CustomerId      string `json:"customerId"`
// 	CustomerName    string `json: "custName"`
// 	CustomerAddress string `json: "custAddress"`
// }

type CreateCustomer struct {
	CustomerName    string `json: "custName"`
	CustomerAddress string `json: "custAddress"`
}
