package usecase

import (
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"invoiceBuana/repository"
	"invoiceBuana/utils"
)

type CustomerUsecase interface {
	UpdateCustomer(customerName string, customerID *string) error
	GetAllCustomer(limit, offset string) ([]dto.DisplayCustomer, error)
	CreateCustomer(request dto.CreateCustomer) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func (r *customerUsecase) UpdateCustomer(customerName string, customerID *string) error {
	customer, err := r.customerRepo.GetCustomerByName(customerName)
	if err != nil {
		return err
	}

	if customer == nil {
		newCustomerID, err := r.customerRepo.CreateCustomer(customerName)
		if err != nil {
			return err
		}
		*customerID = newCustomerID
	} else {
		*customerID = customer.CustomerID
	}

	return nil
}

func (r *customerUsecase) GetAllCustomer(limit, offset string) ([]dto.DisplayCustomer, error) {
	dataCustomer, err := r.customerRepo.GetAll(limit, offset)

	if err != nil {
		return []dto.DisplayCustomer{}, utils.GetCustomerError()
	}

	return dataCustomer, nil
}

func (r *customerUsecase) CreateCustomer(request dto.CreateCustomer) error {

	// generate customer id

	generateCustomerId, err := utils.GenerateUserID()

	if err != nil {
		return utils.CreateIdError()
	}

	// check for duplicate

	dataDuplicate, err := r.customerRepo.GetDuplicateByName(request.CustomerName)

	if err != nil {
		return utils.CreateCustomerError()
	}

	if dataDuplicate == true {
		return utils.DuplicateCustomer()
	}

	// insert customer

	dataCustomer := &model.Customer{
		CustomerID:      generateCustomerId,
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
	}

	errInsert := r.customerRepo.Create(dataCustomer)

	if errInsert != nil {
		return utils.CreateCustomerError()
	}

	return nil

}

func NewCustomerUsecase(customerRepo repository.CustomerRepository) CustomerUsecase {
	usecase := new(customerUsecase)
	usecase.customerRepo = customerRepo
	return usecase
}
