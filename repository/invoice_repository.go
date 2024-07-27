package repository

import (
	"fmt"
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"time"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	UpdateInvoiceDetails(invoiceID string, req dto.UpdateInvoiceRequest, customerID string) error
	DeleteInvoiceItems(invoiceID string) error

	GetInvoiceAll(limit, offset string) ([]dto.InvoiceDetailResponse, error)
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

func (r *invoiceRepository) UpdateInvoiceDetails(invoiceID string, req dto.UpdateInvoiceRequest, customerID string) error {
	query := `UPDATE Invoice 
              SET subject = ?, issue_date = ?, due_date = ?, customer_id = ? 
              WHERE invoice_id = ?`
	result := r.db.Exec(query, req.Subject, req.IssueDate, req.DueDate, customerID, invoiceID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *invoiceRepository) DeleteInvoiceItems(invoiceID string) error {
	query := `DELETE FROM InvoiceItem WHERE invoice_id = ?`
	err := r.db.Exec(query, invoiceID).Error
	return err
}

// GetInvoiceAll retrieves all invoices with pagination
func (r *invoiceRepository) GetInvoiceAll(limit, offset string) ([]dto.InvoiceDetailResponse, error) {
	var invoiceDetails []dto.InvoiceDetailResponse

	// Define the query to retrieve all invoices with related customer name and item count with pagination
	query := `
		SELECT i.invoice_id, i.issue_date, i.subject, COUNT(ii.id) AS total_item, c.customer_name, i.due_date, i.status
		FROM Invoice i
		LEFT JOIN InvoiceItem ii ON i.invoice_id = ii.invoice_id
		LEFT JOIN Customer c ON i.customer_id = c.customer_id
		GROUP BY i.invoice_id, i.issue_date, i.subject, c.customer_name, i.due_date, i.status
		ORDER BY i.invoice_id
		LIMIT ? OFFSET ?
	`

	// Execute the query and scan the result into the slice of InvoiceDetailResponse
	rows, err := r.db.Raw(query, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var invoiceDetail dto.InvoiceDetailResponse
		err := rows.Scan(&invoiceDetail.InvoiceID, &invoiceDetail.IssueDate, &invoiceDetail.SubjectInvoice, &invoiceDetail.TotalItem, &invoiceDetail.CustomerName, &invoiceDetail.DueDate, &invoiceDetail.Status)
		if err != nil {
			return nil, err
		}
		invoiceDetails = append(invoiceDetails, invoiceDetail)
	}

	return invoiceDetails, nil
}

// Get Invoice by ID with related data
func (r *invoiceRepository) GetInvoiceByID(invoiceID string) (model.Invoice, error) {
	var invoice model.Invoice
	var customer model.Customer
	var items []model.InvoiceItem
	// var itemDetail model.Item

	// Retrieve the main invoice
	query := `SELECT * FROM Invoice WHERE invoice_id = ?`
	err := r.db.Raw(query, invoiceID).Scan(&invoice).Error
	if err != nil {
		return invoice, err
	}

	// Retrieve the customer associated with the invoice
	customerQuery := `SELECT * FROM Customer WHERE customer_id = ?`
	err = r.db.Raw(customerQuery, invoice.CustomerID).Scan(&customer).Error
	if err != nil {
		return invoice, err
	}
	invoice.Customer = customer

	// Retrieve the items associated with the invoice
	// itemsQuery := `SELECT ii.*, i.*
	//                FROM InvoiceItem ii
	//                JOIN Item i ON ii.item_id = i.item_id
	//                WHERE ii.invoice_id = ?`
	// err = r.db.Raw(itemsQuery, invoiceID).Scan(&items).Error
	// if err != nil {
	// 	return invoice, err
	// }
	// invoice.Items = items

	// return invoice, nil

	// Retrieve the items associated with the invoice
	itemsQuery := `SELECT ii.id, ii.invoice_id, ii.item_id, ii.quantity, ii.total_price, i.id, i.item_id, i.item_name, i.item_type, i.item_price, i.created_at, i.updated_at
                   FROM InvoiceItem ii 
                   JOIN Item i ON ii.item_id = i.item_id 
                   WHERE ii.invoice_id = ?`
	rows, err := r.db.Raw(itemsQuery, invoiceID).Rows()
	if err != nil {
		return invoice, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.InvoiceItem
		var itemDetail model.Item
		if err := rows.Scan(
			&item.ID, &item.InvoiceID, &item.ItemID, &item.Quantity, &item.TotalPrice,
			&itemDetail.ID, &itemDetail.Item_ID, &itemDetail.Item_Name, &itemDetail.Item_Type, &itemDetail.Item_Price, &itemDetail.CreatedAt, &itemDetail.UpdatedAt,
		); err != nil {
			return invoice, err
		}
		item.Item = itemDetail
		items = append(items, item)
	}
	fmt.Println("ini isi items -==> ", items)

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

		totalPrice := float64(itemReq.Quantity) * item.Item_Price

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
