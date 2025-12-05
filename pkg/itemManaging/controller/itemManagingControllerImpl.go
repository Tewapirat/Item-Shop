package controller

import (
	"net/http"

	"github.com/TewApirat/items-shop-api/pkg/custom"
	_itemManagingModel "github.com/TewApirat/items-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/TewApirat/items-shop-api/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type itemManagingControllerImpl struct {

itemManagingService _itemManagingService.ItemManagingService

}

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController{
	return &itemManagingControllerImpl{itemManagingService}
}

func (c * itemManagingControllerImpl)Creating(pctx echo.Context) error{
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customeEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customeEchoRequest.Bind(itemCreatingReq); err != nil{
	return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil{
		return custom.Error(pctx,http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusCreated, item)

}