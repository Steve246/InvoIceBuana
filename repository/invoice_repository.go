package repository

import (
	"fmt"
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"time"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	GetInvoiceByID(invoiceID string) (model.Invoice, error)
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	InsertInvoice(invoiceId string, invoice dto.InvoiceRequest) (string, error)
	InsertInvoiceItems(invoiceID string, items []dto.InvoiceItemRequest) (float64, error)
	UpdateInvoiceTotals(invoiceID string, subTotal float64) error
}

type invoiceRepository struct {
	db *gorm.DB
}

// Get Invoice by ID with related data
func (r *invoiceRepository) GetInvoiceByID(invoiceID string) (model.Invoice, error) {
	var invoice model.Invoice
	var customer model.Customer
	var items []model.InvoiceItem

	// Retrieve the main invoice
	query := `SELECT * FROM Invoice WHERE invoice_id = ?`
	err := r.db.Raw(query, invoiceID).Scan(&invoice).Error
	if err != nil {
		return invoice, err
	}

	// Retrieve the customer associated with the invoice
	customerQuery := `SELECT * FROM Customer WHERE id = ?`
	err = r.db.Raw(customerQuery, invoice.CustomerID).Scan(&customer).Error
	if err != nil {
		return invoice, err
	}
	invoice.Customer = customer

	// Retrieve the items associated with the invoice
	itemsQuery := `SELECT ii.*, i.* 
                   FROM InvoiceItem ii 
                   JOIN Item i ON ii.item_id = i.item_id 
                   WHERE ii.invoice_id = ?`
	err = r.db.Raw(itemsQuery, invoiceID).Scan(&items).Error
	if err != nil {
		return invoice, err
	}
	invoice.Items = items

	return invoice, nil
}

func (r *invoiceRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *invoiceRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *invoiceRepository) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *invoiceRepository) InsertInvoice(invoiceId string, invoice dto.InvoiceRequest) (string, error) {
	query := `INSERT INTO Invoice (invoice_id, subject, customer_id, issue_date, due_date, status, created_at, updated_at, sub_total, tax, grand_total) 
			  VALUES (?, ?, ?, ?, ?, 'unpaid', ?, ?, 0.00, 0.00, 0.00)`
	result := r.db.Exec(query, invoiceId, invoice.Subject, invoice.CustomerID, invoice.IssueDate, invoice.DueDate, time.Now(), time.Now())
	if result.Error != nil {
		return "", result.Error
	}

	// var invoiceID string
	// r.db.Raw("SELECT LAST_INSERT_ID()").Scan(&invoiceID)
	// if invoiceID == 0 {
	// 	return 0, fmt.Errorf("failed to retrieve last insert ID")
	// }

	return invoiceId, nil
}

func (r *invoiceRepository) InsertInvoiceItems(invoiceID string, items []dto.InvoiceItemRequest) (float64, error) {

	var subTotal float64
	for _, itemReq := range items {
		var item model.Item
		err := r.db.Raw("SELECT * FROM Item WHERE item_id = ?", itemReq.ItemID).Scan(&item).Error
		if err != nil {
			return 0, err
		}

		fmt.Println("ini ada item quantity ==>", itemReq.Quantity)

		fmt.Println("ini ada item price ==>", item.Item_ID)

		totalPrice := float64(itemReq.Quantity) * item.Item_Price
		fmt.Println("ini totalPrice ==> ", totalPrice)
		subTotal += totalPrice

		query := `INSERT INTO InvoiceItem (invoice_id, item_id, quantity, total_price) 
				 VALUES (?, ?, ?, ?)`
		result := r.db.Exec(query, invoiceID, itemReq.ItemID, itemReq.Quantity, totalPrice)
		if result.Error != nil {
			return 0, result.Error
		}
	}

	return subTotal, nil
}

func (r *invoiceRepository) UpdateInvoiceTotals(invoiceID string, subTotal float64) error {
	tax := subTotal * 0.1
	grandTotal := subTotal + tax

	query := `UPDATE Invoice SET sub_total = ?, tax = ?, grand_total = ? WHERE invoice_id = ?`
	result := r.db.Exec(query, subTotal, tax, grandTotal, invoiceID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	repo := new(invoiceRepository)
	repo.db = db
	return repo
}
