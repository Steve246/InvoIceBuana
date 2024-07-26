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

	routerDev.POST("/add/invoice", controller.addNewInvoice)
}
