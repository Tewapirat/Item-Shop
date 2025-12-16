package repository

import (
	databases "github.com/TewApirat/items-shop-api/databases"
	entities "github.com/TewApirat/items-shop-api/entities"
	_inventoryException "github.com/TewApirat/items-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
)


type inventoryRepositoryImpl struct {
	db databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger)InventoryRepository{
	return &inventoryRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r * inventoryRepositoryImpl)Filling(inventoryEntities []*entities.Inventory)([]*entities.Inventory, error){
	inventoryEntitiesResult := make([]*entities.Inventory, 0)

	if err := r.db.Connect().CreateInBatches(
		inventoryEntities, 
		len(inventoryEntities),
	).Scan(&inventoryEntitiesResult).Error; err != nil {
		r.logger.Errorf("error filling inventory: %s",err.Error())
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: inventoryEntities[0].PlayerID,
			ItemID: inventoryEntities[0].ItemID,
		}
	}

	return inventoryEntitiesResult, nil
}

func (r *inventoryRepositoryImpl)Removing(playerID string, itemID uint64, limit int)error {

	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	tx := r.db.Connect().Begin()


	for _, inventory := range inventoryEntities{
		inventory.IsDeleted = true

		if err := tx.Model(&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s",err.Error())
			return &_inventoryException.PlayerItemRemoving{
				ItemID: itemID,
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		r.logger.Errorf("error removing player item in inventory: %s",err.Error())
		return &_inventoryException.PlayerItemRemoving{
			ItemID: itemID,
		}
	}
		
	return nil
}


func (r *inventoryRepositoryImpl)PlayerItemCounting(playerId string, itemID uint64) int64{
	var count int64

	if err := r.db.Connect().Model(
		&entities.Inventory{},
	).Where(
		"player_id = ? and item_id = ? and is_deleted = ?",playerId, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory: %s", err.Error())
		return -1
	}
	return count
}

func(r *inventoryRepositoryImpl)Listing(playerID string)([]*entities.Inventory, error){
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and is_deleted = ? ", playerID, false,
	).Find(&inventoryEntities).Error; err != nil{
		r.logger.Errorf("error listing player inventory: %s",err.Error())
		return nil, &_inventoryException.PlayerItemsFinding{
			PlayerID: playerID,
		}
	}
	
	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(
	playerID string,
	itemID uint64,
	limit int,
) ([]*entities.Inventory, error){
	inventoryEntities := make([]*entities.Inventory,0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventoryEntities).Error; err != nil {
			r.logger.Errorf("error finding player item in inventory by ID: %s", err.Error())
			return nil, &_inventoryException.PlayerItemRemoving{
				ItemID: itemID,
			}

	}

	return inventoryEntities, nil
}