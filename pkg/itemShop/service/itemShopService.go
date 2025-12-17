package service

import (
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFiler *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq)(*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.BuyingReq)(*_playerCoinModel.PlayerCoin, error)
}