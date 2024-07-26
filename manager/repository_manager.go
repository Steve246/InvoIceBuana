package manager

import "invoiceBuana/repository"

type RepositoryManager interface {
	InvoiceRepo() repository.InvoiceRepository
	CustomerRepo() repository.CustomerRepository
	ItemRepo() repository.ItemRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) InvoiceRepo() repository.InvoiceRepository {
	return repository.NewInvoiceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ItemRepo() repository.ItemRepository {
	return repository.NewItemRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
