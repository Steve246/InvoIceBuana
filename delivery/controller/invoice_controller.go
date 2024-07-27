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

func (i *InvoiceController) getInvoiceById(c *gin.Context) {
	idParam := c.Query("id")

	data, err := i.ucInvoice.GetInvoiceByID(idParam)

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

	routerDev.GET("/display/customInvoice", controller.getInvoiceById)

	routerDev.GET("/display/invoice", controller.getAllInvoice)

	routerDev.POST("/add/invoice", controller.addNewInvoice)
}
