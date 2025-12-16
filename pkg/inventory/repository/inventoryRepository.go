package repository

import "github.com/TewApirat/items-shop-api/entities"

type InventoryRepository interface {
	Filling(inventoryEntities []*entities.Inventory)([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int)error
	PlayerItemCounting(playerId string, itemID uint64) int64
	Listing(playerId string)([]*entities.Inventory, error)

}