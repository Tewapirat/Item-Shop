package service

import (
	_inventoryModel "github.com/TewApirat/items-shop-api/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string)([]*_inventoryModel.Inventory, error)
}