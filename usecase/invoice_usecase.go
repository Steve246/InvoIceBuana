package usecase

import (
	"fmt"
	"invoiceBuana/model/dto"
	"invoiceBuana/repository"
	"invoiceBuana/utils"
)

type InvoiceUsecase interface {
	UpdateInvoice(invoiceID string, req dto.UpdateInvoiceRequest) error
	GetInvoiceByID(InvoiceId string) (dto.InvoiceResponse, error)
	GetInvoiceAll(limit, offset string) ([]dto.InvoiceDetailResponse, error)
	CreateInvoice(invoice dto.InvoiceRequest) (dto.InvoiceResponse, error)
}

type invoiceUsecase struct {
	itemRepo     repository.ItemRepository
	customerRepo repository.CustomerRepository
	invoiceRepo  repository.InvoiceRepository
	invoiceUtils utils.InvoiceCounter
}

func (u *invoiceUsecase) UpdateInvoice(invoiceID string, req dto.UpdateInvoiceRequest) error {
	tx := u.invoiceRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			u.invoiceRepo.Rollback(tx)
		}
	}()

	// pastiin id customer yg diinput, itu ada
	_, err := u.customerRepo.GetById(req.CustomerID)
	if err != nil {
		u.invoiceRepo.Rollback(tx)
		return fmt.Errorf("customer with ID %s does not exist", req.CustomerID)
	}

	// update data invoice, dengan customer id terbaru
	err = u.invoiceRepo.UpdateInvoiceDetails(invoiceID, req, req.CustomerID)
	if err != nil {
		u.invoiceRepo.Rollback(tx)
		return err
	}

	// kalau ada item yg di-delete user, hilangin bersamaan dgn relasinya
	err = u.invoiceRepo.DeleteInvoiceItems(invoiceID)
	if err != nil {
		u.invoiceRepo.Rollback(tx)
		return err
	}

	// masukin item terbaru
	_, err = u.invoiceRepo.InsertInvoiceItems(invoiceID, req.Items)
	if err != nil {
		u.invoiceRepo.Rollback(tx)
		return err
	}

	// Commit transaction
	err = u.invoiceRepo.Commit(tx)
	if err != nil {
		return err
	}

	return nil
}

func (u *invoiceUsecase) GetInvoiceByID(InvoiceId string) (dto.InvoiceResponse, error) {
	// Retrieve the created invoice with all related data

	createdInvoice, err := u.invoiceRepo.GetInvoiceByID(InvoiceId)
	if err != nil {
		return dto.InvoiceResponse{}, err
	}

	// retrieve item

	response := dto.InvoiceResponse{
		ID:         createdInvoice.ID,
		InvoiceID:  createdInvoice.InvoiceID,
		Subject:    createdInvoice.Subject,
		CustomerID: createdInvoice.CustomerID,
		IssueDate:  createdInvoice.IssueDate.Format("2006-01-02"),
		DueDate:    createdInvoice.DueDate.Format("2006-01-02"),
		Status:     createdInvoice.Status,
		Customer: dto.CustomerResponse{
			// ID:      customerData.ID,
			// Name:    customerData.CustomerName,
			// Address: customerData.CustomerAddress,
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

		// fmt.Println("ini dapet Item_ID ==>", item.ItemID)

		// dataItem, err := u.itemRepo.GetById(item.ItemID)

		if err != nil {
			// fmt.Println("ini error ==> ", err)
			return dto.InvoiceResponse{}, err
		}

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

func (u *invoiceUsecase) GetInvoiceAll(limit, offset string) ([]dto.InvoiceDetailResponse, error) {
	dataInvoices, err := u.invoiceRepo.GetInvoiceAll(limit, offset)
	if err != nil {
		return nil, utils.GetInvoiceError()
	}

	var dataInvoiceMapping []dto.InvoiceDetailResponse
	for _, dataInvoice := range dataInvoices {
		dataInvoiceMapping = append(dataInvoiceMapping, dto.InvoiceDetailResponse{
			InvoiceID:      dataInvoice.InvoiceID,
			IssueDate:      dataInvoice.IssueDate,
			SubjectInvoice: dataInvoice.SubjectInvoice,
			TotalItem:      dataInvoice.TotalItem,
			CustomerName:   dataInvoice.CustomerName,
			DueDate:        dataInvoice.DueDate,
			Status:         dataInvoice.Status,
		})
	}

	return dataInvoiceMapping, nil
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

	// retrieve item

	response := dto.InvoiceResponse{
		ID:         createdInvoice.ID,
		InvoiceID:  createdInvoice.InvoiceID,
		Subject:    createdInvoice.Subject,
		CustomerID: createdInvoice.CustomerID,
		IssueDate:  createdInvoice.IssueDate.Format("2006-01-02"),
		DueDate:    createdInvoice.DueDate.Format("2006-01-02"),
		Status:     createdInvoice.Status,
		Customer: dto.CustomerResponse{
			// ID:      customerData.ID,
			// Name:    customerData.CustomerName,
			// Address: customerData.CustomerAddress,
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

		// fmt.Println("ini dapet Item_ID ==>", item.ItemID)

		// dataItem, err := u.itemRepo.GetById(item.ItemID)

		if err != nil {
			// fmt.Println("ini error ==> ", err)
			return dto.InvoiceResponse{}, err
		}

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

func NewInvoiceUsecase(invoiceRepo repository.InvoiceRepository, invoiceUtils utils.InvoiceCounter, itemRepo repository.ItemRepository, customerRepo repository.CustomerRepository) InvoiceUsecase {
	usecase := new(invoiceUsecase)
	usecase.invoiceRepo = invoiceRepo
	usecase.invoiceUtils = invoiceUtils
	usecase.customerRepo = customerRepo
	usecase.itemRepo = itemRepo
	return usecase
}
