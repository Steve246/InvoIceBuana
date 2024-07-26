package usecase

import (
	"fmt"
	"invoiceBuana/model/dto"
	"invoiceBuana/repository"
	"invoiceBuana/utils"
)

type InvoiceUsecase interface {
	CreateInvoice(invoice dto.InvoiceRequest) (dto.InvoiceResponse, error)
}

type invoiceUsecase struct {
	invoiceRepo  repository.InvoiceRepository
	invoiceUtils utils.InvoiceCounter
}

func (u *invoiceUsecase) CreateInvoice(invoice dto.InvoiceRequest) (dto.InvoiceResponse, error) {

	tx := u.invoiceRepo.Begin()
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return dto.InvoiceResponse{}, tx.Error
	}

	// Initialize the counter table (it should be done once, probably not needed here if already initialized)
	err := u.invoiceUtils.InitializeCounterTable()
	if err != nil {
		tx.Rollback()
		return dto.InvoiceResponse{}, utils.CreateIdError()
	}

	// Generate invoice ID
	generateInvId, err := u.invoiceUtils.GenerateInvoiceID()
	if err != nil {
		tx.Rollback()
		return dto.InvoiceResponse{}, utils.CreateIdError()
	}

	// Insert Invoice
	invoiceID, err := u.invoiceRepo.InsertInvoice(generateInvId, invoice)

	if err != nil {
		tx.Rollback()
		return dto.InvoiceResponse{}, err
	}

	// Insert Invoice Items and calculate Subtotal
	subTotal, err := u.invoiceRepo.InsertInvoiceItems(invoiceID, invoice.Items)
	if err != nil {
		tx.Rollback()
		return dto.InvoiceResponse{}, err
	}

	fmt.Println("ini subTotal ==> ", subTotal)

	// Update Invoice with calculated totals
	err = u.invoiceRepo.UpdateInvoiceTotals(invoiceID, subTotal)
	if err != nil {
		tx.Rollback()
		return dto.InvoiceResponse{}, err
	}

	tx.Commit()

	// Retrieve the created invoice with all related data
	createdInvoice, err := u.invoiceRepo.GetInvoiceByID(invoiceID)
	if err != nil {
		return dto.InvoiceResponse{}, err
	}

	fmt.Println("ini data invoice ===>", createdInvoice)

	response := dto.InvoiceResponse{
		ID:         createdInvoice.ID,
		InvoiceID:  createdInvoice.InvoiceID,
		Subject:    createdInvoice.Subject,
		CustomerID: createdInvoice.CustomerID,
		IssueDate:  createdInvoice.IssueDate.Format("2006-01-02"),
		DueDate:    createdInvoice.DueDate.Format("2006-01-02"),
		Status:     createdInvoice.Status,
		Customer: dto.CustomerResponse{
			ID:      createdInvoice.Customer.ID,
			Name:    createdInvoice.Customer.CustomerName,
			Address: createdInvoice.Customer.CustomerAddress,
		},
		Items: []dto.InvoiceItemResponse{},
		Totals: dto.TotalsResponse{
			TotalItems: len(createdInvoice.Items),
			Subtotal:   createdInvoice.SubTotal,
			Tax:        createdInvoice.Tax,
			GrandTotal: createdInvoice.GrandTotal,
		},
		CreatedAt: createdInvoice.CreatedAt,
		UpdatedAt: createdInvoice.UpdatedAt,
	}

	for _, item := range createdInvoice.Items {
		response.Items = append(response.Items, dto.InvoiceItemResponse{
			ID:         item.ID,
			ItemName:   item.Item.Item_Name,
			Quantity:   item.Quantity,
			UnitPrice:  item.Item.Item_Price,
			TotalPrice: item.TotalPrice,
		})
	}

	return response, nil
}

func NewInvoiceUsecase(invoiceRepo repository.InvoiceRepository, invoiceUtils utils.InvoiceCounter) InvoiceUsecase {
	usecase := new(invoiceUsecase)
	usecase.invoiceRepo = invoiceRepo
	usecase.invoiceUtils = invoiceUtils
	return usecase
}