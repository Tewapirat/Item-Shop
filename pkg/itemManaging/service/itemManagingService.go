package service

import (
	_itemManagingModel "github.com/TewApirat/items-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(ItemCreatingReq *_itemManagingModel.ItemCreatingReq) ( *_itemShopModel.Item,error)
	Editing(itemID uint64, itemEditing * _itemManagingModel.ItemEditingReq)(*_itemShopModel.Item, error)
}