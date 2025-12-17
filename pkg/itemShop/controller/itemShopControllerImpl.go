package controller

import (
	"net/http"

	"github.com/TewApirat/items-shop-api/pkg/custom"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/TewApirat/items-shop-api/pkg/itemShop/service"
	"github.com/TewApirat/items-shop-api/pkg/playerCoin/validation"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}


func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {

	itemFilter := new(_itemShopModel.ItemFilter)

	customeEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customeEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemModelList)

}

func (c *itemShopControllerImpl)Buying(pctx echo.Context)error{
	
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest,err)
	}
	
	buyingReq := new(_itemShopModel.BuyingReq)

	customeEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customeEchoRequest.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}


	buyingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil{
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	
	return pctx.JSON(http.StatusOK,playerCoin)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	sellingReq := new(_itemShopModel.SellingReq)

	validatingContext := custom.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(sellingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID

	result, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}
