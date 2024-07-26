package utils

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// FIXME: need refactor error code

type InvoiceCounter interface {
	GenerateInvoiceID() (string, error)
	InitializeCounterTable() error
}

type invoiceCounter struct {
	db *gorm.DB
	mu sync.Mutex
}

// GenerateInvoiceID generates the next invoice ID in the sequence
func (ic *invoiceCounter) GenerateInvoiceID() (string, error) {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	var lastID int
	// Using GORM to query the last invoice ID
	err := ic.db.Raw("SELECT last_invoice_id FROM InvoiceSequence WHERE id = ?", 1).Scan(&lastID).Error
	if err != nil {

		return "", fmt.Errorf("failed to query last invoice ID: %w", err)
	}

	nextID := lastID + 1
	// Using GORM to update the invoice ID
	err = ic.db.Exec("UPDATE InvoiceSequence SET last_invoice_id = ? WHERE id = ?", nextID, 1).Error
	if err != nil {
		return "", fmt.Errorf("failed to update invoice ID: %w", err)
	}

	return fmt.Sprintf("INV%03d", nextID), nil
}

// InitializeCounterTable ensures the InvoiceSequence table exists and has an initial value
func (ic *invoiceCounter) InitializeCounterTable() error {
	// Create the table if it does not exist

	fmt.Println("masuk sini")
	err := ic.db.Exec(`
		CREATE TABLE IF NOT EXISTS InvoiceSequence (
			id INT PRIMARY KEY,
			last_invoice_id INT
		);
	`).Error
	if err != nil {
		return fmt.Errorf("failed to create InvoiceSequence table: %w", err)
	}

	// Initialize the row if it does not exist
	var count int
	err = ic.db.Raw("SELECT COUNT(*) FROM InvoiceSequence WHERE id = ?", 1).Scan(&count).Error
	if err != nil {
		return fmt.Errorf("failed to check InvoiceSequence row: %w", err)
	}

	if count == 0 {
		err = ic.db.Exec("INSERT INTO InvoiceSequence (id, last_invoice_id) VALUES (?, ?)", 1, 0).Error
		if err != nil {
			return fmt.Errorf("failed to initialize InvoiceSequence row: %w", err)
		}
	}

	return nil
}

// NewInvoiceCounter initializes and returns a new InvoiceCounter
func NewInvoiceCounter(db *gorm.DB) InvoiceCounter {
	return &invoiceCounter{
		db: db,
	}
}
