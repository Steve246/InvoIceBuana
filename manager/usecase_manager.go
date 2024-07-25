package manager

import "invoiceBuana/usecase"

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
	ItemUsecase() usecase.ItemUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) ItemUsecase() usecase.ItemUsecase {
	return usecase.NewItemUsecase(u.repoManager.ItemRepo())
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
