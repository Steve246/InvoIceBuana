package usecase

import (
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"invoiceBuana/repository"
	"invoiceBuana/utils"

	"gorm.io/gorm"
)

type CustomerUsecase interface {
	GetAllCustomer() ([]model.Customer, error)
	CreateCustomer(request dto.CreateCustomer) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func (r *customerUsecase) GetAllCustomer() ([]model.Customer, error) {
	dataCustomer, err := r.customerRepo.GetAll()

	if err != nil {
		return []model.Customer{}, err
	}

	return dataCustomer, nil
}

func (r *customerUsecase) CreateCustomer(request dto.CreateCustomer) error {

	generateCustomerId, err := utils.GenerateUserID()

	if err != nil {
		return utils.IdCustomerError()
	}

	dataCustomer := &model.Customer{
		Model:           gorm.Model{},
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
