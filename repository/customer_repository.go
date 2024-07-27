package repository

import (
	"fmt"
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"invoiceBuana/utils"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	UpdateCustomer(customerID string, name string) error

	CreateCustomer(name string) (string, error)
	GetCustomerByName(name string) (*model.Customer, error)

	GetById(customer_id string) (model.Customer, error)
	GetDuplicateByName(name string) (bool, error)
	GetAll(limit, offset string) ([]dto.DisplayCustomer, error)
	Create(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

// UpdateCustomer updates the customer name based on the customer ID
func (r *customerRepository) UpdateCustomer(customerID string, name string) error {
	query := `UPDATE Customer SET customer_name = ? WHERE customer_id = ?`
	result := r.db.Exec(query, name, customerID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *customerRepository) CreateCustomer(name string) (string, error) {
	customerID, _ := utils.GenerateUserID() // Generate a new ID as per your logic
	query := `INSERT INTO Customer (customer_id, customer_name) VALUES (?, ?)`
	result := r.db.Exec(query, customerID, name)
	if result.Error != nil {
		return "", result.Error
	}
	return customerID, nil
}

func (r *customerRepository) GetCustomerByName(name string) (*model.Customer, error) {
	var customer model.Customer
	query := `SELECT * FROM Customer WHERE customer_name = ?`
	err := r.db.Raw(query, name).Scan(&customer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetById(customer_id string) (model.Customer, error) {
	var customer model.Customer

	query := `SELECT * from Customer WHERE customer_id = ?`

	err := r.db.Raw(query, customer_id).Scan(&customer).Error
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
