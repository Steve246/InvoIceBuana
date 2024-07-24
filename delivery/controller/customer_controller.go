package controller

import (
	"invoiceBuana/delivery/api"
	"invoiceBuana/model/dto"
	"invoiceBuana/usecase"
	"invoiceBuana/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	routerDev  *gin.RouterGroup
	ucCustomer usecase.CustomerUsecase
	api.BaseApi
}

func (u *CustomerController) addNewCustomer(c *gin.Context) {
	var bodyRequest dto.CreateCustomer

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	err := u.ucCustomer.CreateCustomer(bodyRequest)
	if err != nil {
		u.Failed(c, err)
		return

	}

	detailMsg := "Customer Created Succesfully"
	u.Success(c, "", detailMsg, "register")
}

func NewCustomerController(routerDev *gin.RouterGroup, ucCustomer usecase.CustomerUsecase) {
	controller := CustomerController{
		routerDev:  routerDev,
		ucCustomer: ucCustomer,
		BaseApi:    api.BaseApi{},
	}

	routerDev.POST("/add/items", controller.addNewCustomer)
}
