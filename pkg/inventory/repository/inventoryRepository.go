package repository

import (
	"github.com/TewApirat/items-shop-api/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int)([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, limit int)error
	PlayerItemCounting(playerId string, itemID uint64) int64
	Listing(playerId string)([]*entities.Inventory, error)

}