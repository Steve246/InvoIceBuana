package dto

import "time"

// for updated invoice

type UpdateInvoiceRequest struct {
	Subject    string               `json:"subject"`
	IssueDate  string               `json:"issue_date"`
	DueDate    string               `json:"due_date"`
	CustomerID string               `json:"customer_id"`
	Items      []InvoiceItemRequest `json:"items"`
}

// type InvoiceItemUpdateRequest struct {
// 	ItemID   string `json:"item_id"`
// 	Quantity int    `json:"quantity"`
// }

type InvoiceDetailResponse struct {
	InvoiceID      string `json:"invoice_id"`
	IssueDate      string `json:"issue_date"`
	SubjectInvoice string `json:"subject_invoice"`
	TotalItem      int    `json:"total_item"`
	CustomerName   string `json:"customer_name"`
	DueDate        string `json:"due_date"`
	Status         string `json:"status"`
}

type InvoiceRequest struct {
	// InvoiceID  string               `json:"invoice_id"`
	Subject    string               `json:"subject"`
	CustomerID string               `json:"customer_id"`
	IssueDate  string               `json:"issue_date"`
	DueDate    string               `json:"due_date"`
	Items      []InvoiceItemRequest `json:"items"`
}

type InvoiceItemRequest struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type InvoiceResponse struct {
	ID         uint                  `json:"id"`
	InvoiceID  string                `json:"invoice_id"`
	Subject    string                `json:"subject"`
	CustomerID string                `json:"customer_id"`
	IssueDate  string                `json:"issue_date"`
	DueDate    string                `json:"due_date"`
	Status     string                `json:"status"`
	Customer   CustomerResponse      `json:"customer"`
	Items      []InvoiceItemResponse `json:"items"`
	Totals     TotalsResponse        `json:"totals"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

type InvoiceItemResponse struct {
	ID         uint    `json:"id"`
	ItemName   string  `json:"item_name"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}

type CustomerResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type TotalsResponse struct {
	TotalItems int     `json:"total_items"`
	Subtotal   float64 `json:"subtotal"`
	Tax        float64 `json:"tax"`
	GrandTotal float64 `json:"grand_total"`
}
