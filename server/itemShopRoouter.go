package server

import (
	_itemShopRepository "github.com/TewApirat/items-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/TewApirat/items-shop-api/pkg/itemShop/service"
	_itemShopController "github.com/TewApirat/items-shop-api/pkg/itemShop/controller"
	_inventoryRepository "github.com/TewApirat/items-shop-api/pkg/inventory/repository"
	_playerCoinRepository "github.com/TewApirat/items-shop-api/pkg/playerCoin/repository"

)

func (s *echoServer)initItemShopRouter(m *authorizingMiddleware){
	router := s.app.Group("v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryReository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)


	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		playerCoinRepository,
		inventoryReository,
		s.app.Logger,
	)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("",itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)

}