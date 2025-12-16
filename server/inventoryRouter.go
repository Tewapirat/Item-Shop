package server

import (
	_inventoryRepository "github.com/TewApirat/items-shop-api/pkg/inventory/repository"
	_inventoryService "github.com/TewApirat/items-shop-api/pkg/inventory/service"
	_inventoryController "github.com/TewApirat/items-shop-api/pkg/inventory/controller"
	_itemShopRepository "github.com/TewApirat/items-shop-api/pkg/itemShop/repository"
)
func (s *echoServer) initinventoryRouter(m *authorizingMiddleware){
	router := s.app.Group("/v1/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(
		inventoryRepository,
		itemShopRepository,
	)
	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)


	router.GET("", inventoryController.Listing,m.PlayerAuthorizing)

}