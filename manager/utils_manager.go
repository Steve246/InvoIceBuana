package manager

import "invoiceBuana/utils"

type UtilsManager interface {
	InvoiceIdUtils() utils.InvoiceCounter
}

type utilsManager struct {
	infra Infra
}

func (u *utilsManager) InvoiceIdUtils() utils.InvoiceCounter {
	return utils.NewInvoiceCounter(u.infra.SqlDb())
}

func NewUtilsManager(infra Infra) UtilsManager {
	return &utilsManager{
		infra: infra,
	}
}
