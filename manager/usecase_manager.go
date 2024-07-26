package manager

import "invoiceBuana/usecase"

type UsecaseManager interface {
	InvoiceUsecase() usecase.InvoiceUsecase
	CustomerUsecase() usecase.CustomerUsecase
	ItemUsecase() usecase.ItemUsecase
}

type usecaseManager struct {
	repoManager  RepositoryManager
	utilsManager UtilsManager
}

func (u *usecaseManager) InvoiceUsecase() usecase.InvoiceUsecase {
	return usecase.NewInvoiceUsecase(u.repoManager.InvoiceRepo(), u.utilsManager.InvoiceIdUtils())
}

func (u *usecaseManager) ItemUsecase() usecase.ItemUsecase {
	return usecase.NewItemUsecase(u.repoManager.ItemRepo())
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func NewUsecaseManager(repoManager RepositoryManager, utilsManager UtilsManager) UsecaseManager {
	return &usecaseManager{
		repoManager:  repoManager,
		utilsManager: utilsManager,
	}
}
