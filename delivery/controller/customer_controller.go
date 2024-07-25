package controller

import (
	"fmt"
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

func (u *CustomerController) getAllCustomer(c *gin.Context) {

	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	data, err := u.ucCustomer.GetAllCustomer(limitParam, offsetParam)
	if err != nil {
		u.Failed(c, err)
		return

	}

	detailMsg := "Customer Data Succesfully Retrieved"
	u.Success(c, data, detailMsg, "register")
}

func (u *CustomerController) addNewCustomer(c *gin.Context) {
	var bodyRequest dto.CreateCustomer

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	fmt.Println("ini isi body --> ", bodyRequest)

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

	routerDev.GET("/display/customer", controller.getAllCustomer)

	routerDev.POST("/add/customer", controller.addNewCustomer)
}
