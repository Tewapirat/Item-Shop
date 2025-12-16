package service

import (
	"github.com/TewApirat/items-shop-api/entities"
	_inventoryModel "github.com/TewApirat/items-shop-api/pkg/inventory/model"
	_inventoryRepository "github.com/TewApirat/items-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/TewApirat/items-shop-api/pkg/itemShop/repository"
)


type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(
	inventoryRepository _inventoryRepository.InventoryRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		itemShopRepository: itemShopRepository,

	}
}

func (s *inventoryServiceImpl)Listing(playerID string)([]*_inventoryModel.Inventory, error){
	inventoryEntities, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuanityCounterList := s.getUniqueItemWithQuanityCounterList(inventoryEntities)


	return s.buildInventoryListingReslt(
		uniqueItemWithQuanityCounterList,
	), nil

}


func (s *inventoryServiceImpl) getUniqueItemWithQuanityCounterList(
	inventoryEntities []*entities.Inventory,
)[]_inventoryModel.ItemQuantityCounting{
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities{
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID: itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList
}


func (s *inventoryServiceImpl)buildInventoryListingReslt(
	uniqueItemWithQuanityCounterList []_inventoryModel.ItemQuantityCounting,
)[]*_inventoryModel.Inventory{

	uniqueItemIDList := s.getItemID(uniqueItemWithQuanityCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory,0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuanityCounterList)


	for _, itemEntity := range itemEntities{
		results = append(results, &_inventoryModel.Inventory{
			Item:		itemEntity.ToItemModel(),
			Quantity: 	itemMapWithQuantity[itemEntity.ID],
			
		})
	}

	return results

}


func (s *inventoryServiceImpl)getItemID(
	uniqueItemWithQuanityCounterList []_inventoryModel.ItemQuantityCounting,
	) []uint64{
		uniqueItemIDList := make([]uint64, 0)

		for _, inventory := range uniqueItemWithQuanityCounterList{
			uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
		}

		return uniqueItemIDList

}


func (s *inventoryServiceImpl)getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
)map[uint64]uint {
	itemMapWithQuantiity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantityCounterList{
		itemMapWithQuantiity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantiity
}