package controller

import (
	"net/http"

	"github.com/TewApirat/items-shop-api/pkg/custom"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/TewApirat/items-shop-api/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService,) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {


	itemFilter := new(_itemShopModel.ItemFilter)

	customeEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customeEchoRequest.Bind(itemFilter); err != nil{
	return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}



	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)

}