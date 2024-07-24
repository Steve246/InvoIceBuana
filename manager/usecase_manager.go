package manager

import "invoiceBuana/usecase"

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
