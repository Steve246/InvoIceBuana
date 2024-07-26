package repository

import (
	"fmt"
	"invoiceBuana/model"
	"invoiceBuana/model/dto"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetById(customer_id string) (model.Customer, error)
	GetDuplicateByName(name string) (bool, error)
	GetAll(limit, offset string) ([]dto.DisplayCustomer, error)
	Create(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func (r *customerRepository) GetById(customer_id string) (model.Customer, error) {
	var customer model.Customer

	query := `SELECT * from Customer WHERE customer_id = ?`

	err := r.db.Raw(query, customer_id).Scan(&customer).Error
	fmt.Println("Ini error ==> ", err)
	if err != nil {

		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) GetDuplicateByName(name string) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM Customer WHERE customer_name = ?"
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

func (r *customerRepository) GetAll(limit, offset string) ([]dto.DisplayCustomer, error) {
	var customers []dto.DisplayCustomer

	query := `SELECT customer_id, customer_name, customer_address FROM Customer LIMIT ? OFFSET ?`

	rows, err := r.db.Raw(query, limit, offset).Rows()
	if err != nil {
		return nil, err
	}

	// closed database resource

	defer rows.Close()

	// remapped using append to get all data customer

	for rows.Next() {
		var customer dto.DisplayCustomer
		if err := rows.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerAddress); err != nil {
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
