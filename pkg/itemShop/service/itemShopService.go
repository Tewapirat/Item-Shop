package service

import (
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFiler *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}