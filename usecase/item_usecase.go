package usecase

import (
	"invoiceBuana/model"
	"invoiceBuana/model/dto"
	"invoiceBuana/repository"
	"invoiceBuana/utils"
)

type ItemUsecase interface {
	GetAllItem(limit, offset string) ([]dto.DisplayItem, error)
	CreateItem(request dto.CreateItem) error
}

type itemUsecase struct {
	itemRepo repository.ItemRepository
}

func (i *itemUsecase) GetAllItem(limit, offset string) ([]dto.DisplayItem, error) {
	dataItem, err := i.itemRepo.GetAll(limit, offset)

	if err != nil {
		return []dto.DisplayItem{}, utils.GetItemError()
	}

	return dataItem, err
}

func (i *itemUsecase) CreateItem(request dto.CreateItem) error {

	// generate customer id

	generateCustomerId, err := utils.GenerateUserID()

	if err != nil {
		return utils.CreateIdError()
	}

	// check duplicate

	dataDuplicate, err := i.itemRepo.GetDuplicateByName(request.ItemName)

	if err != nil {
		return utils.CreateItemsError()
	}

	if dataDuplicate == true {
		return utils.DuplicateItemError()
	}

	// insert customer

	dataItem := &model.Item{
		ItemId:    generateCustomerId,
		ItemName:  request.ItemName,
		ItemType:  request.ItemType,
		ItemPrice: request.ItemPrice,
	}

	errInsert := i.itemRepo.Create(dataItem)

	if errInsert != nil {
		return utils.CreateCustomerError()
	}

	return nil

}

func NewItemUsecase(itemRepo repository.ItemRepository) ItemUsecase {
	usecase := new(itemUsecase)
	usecase.itemRepo = itemRepo
	return usecase
}
