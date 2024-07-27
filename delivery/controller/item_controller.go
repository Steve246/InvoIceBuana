package controller

import (
	"invoiceBuana/delivery/api"
	"invoiceBuana/model/dto"
	"invoiceBuana/usecase"
	"invoiceBuana/utils"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	routerDev *gin.RouterGroup
	ucItem    usecase.ItemUsecase
	api.BaseApi
}

func (i *ItemController) getAllItem(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	data, err := i.ucItem.GetAllItem(limitParam, offsetParam)
	if err != nil {
		i.Failed(c, err)
		return
	}

	detailMsg := "Item Data Succesfully Retrieved"
	i.Success(c, data, detailMsg, "register")

}

func (i *ItemController) addNewItem(c *gin.Context) {
	var bodyRequest dto.CreateItem

	if err := i.ParseRequestBody(c, &bodyRequest); err != nil {
		i.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	err := i.ucItem.CreateItem(bodyRequest)
	if err != nil {
		i.Failed(c, err)
		return

	}

	detailMsg := "Item Created Succesfully"
	i.Success(c, "", detailMsg, "register")
}

func NewItemController(routerDev *gin.RouterGroup, ucItem usecase.ItemUsecase) {
	controller := ItemController{
		routerDev: routerDev,
		ucItem:    ucItem,
		BaseApi:   api.BaseApi{},
	}

	routerDev.GET("/display/items", controller.getAllItem)

	routerDev.POST("/add/items", controller.addNewItem)
}
