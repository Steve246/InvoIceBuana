package repository

import (
	"fmt"
	"invoiceBuana/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetDuplicateByName(name string) (bool, error)
	GetAll() ([]model.Customer, error)
	Create(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func (r *customerRepository) GetDuplicateByName(name string) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM customers WHERE customer_name = ?"
	result := r.db.Raw(query, name).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}

	if count > 0 == true {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *customerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer

	query := `SELECT id, customer_id, customer_name, customer_address, created_at, updated_at FROM Customer`

	rows, err := r.db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	// closed database resource

	defer rows.Close()

	// remapped using append to get all data customer

	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.CustomerID, &customer.CustomerName, &customer.CustomerAddress, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *customerRepository) Create(customer *model.Customer) error {
	query := `INSERT INTO Customer (customer_id, customer_name, customer_address)
	          VALUES (?, ?, ?)`
	result := r.db.Exec(query, customer.CustomerID, customer.CustomerName, customer.CustomerAddress)
	if result.Error != nil {
		return result.Error
	}
	customer.ID = uint(result.RowsAffected)
	fmt.Println(customer.ID)
	return nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
