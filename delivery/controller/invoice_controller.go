package controller

import (
	"invoiceBuana/delivery/api"
	"invoiceBuana/model/dto"
	"invoiceBuana/usecase"
	"invoiceBuana/utils"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	routerDev *gin.RouterGroup
	ucInvoice usecase.InvoiceUsecase
	api.BaseApi
}

func (i *InvoiceController) updateInvoice(c *gin.Context) {
	var bodyRequest dto.UpdateInvoiceRequest

	if err := i.ParseRequestBody(c, &bodyRequest); err != nil {
		i.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	// Extract invoiceID from the URL
	invoiceID := c.Param("id")

	err := i.ucInvoice.UpdateInvoice(invoiceID, bodyRequest)
	if err != nil {
		i.Failed(c, err)
		return
	}

	detailMsg := "Invoice Updated Successfully"
	i.Success(c, nil, detailMsg, "update")
}

func (i *InvoiceController) getInvoiceById(c *gin.Context) {
	invoiceID := c.Param("id")

	data, err := i.ucInvoice.GetInvoiceByID(invoiceID)

	if err != nil {
		i.Failed(c, err)
		return
	}

	detailMsg := "Invoice Data Succesfully Retrieved"
	i.Success(c, data, detailMsg, "register")
}

func (i *InvoiceController) getAllInvoice(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	data, err := i.ucInvoice.GetInvoiceAll(limitParam, offsetParam)

	if err != nil {
		i.Failed(c, err)
		return
	}

	detailMsg := "Invoice Data Succesfully Retrieved"
	i.Success(c, data, detailMsg, "register")
}

func (i *InvoiceController) addNewInvoice(c *gin.Context) {
	var bodyRequest dto.InvoiceRequest

	if err := i.ParseRequestBody(c, &bodyRequest); err != nil {

		i.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, err := i.ucInvoice.CreateInvoice(bodyRequest)
	if err != nil {
		i.Failed(c, err)
		return
	}

	detailMsg := "Invoice Created Sucessfully"
	i.Success(c, data, detailMsg, "register")

}

func NewInvoiceController(routerDev *gin.RouterGroup, ucInvoice usecase.InvoiceUsecase) {
	controller := InvoiceController{
		routerDev: routerDev,
		ucInvoice: ucInvoice,
		BaseApi:   api.BaseApi{},
	}

	routerDev.PUT("/update/invoice/:id", controller.updateInvoice)

	routerDev.GET("/display/customInvoice/:id", controller.getInvoiceById)

	routerDev.GET("/display/invoice", controller.getAllInvoice)

	routerDev.POST("/add/invoice", controller.addNewInvoice)
}
