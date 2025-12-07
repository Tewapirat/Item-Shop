package server

import (
	_itemManagingRepository "github.com/TewApirat/items-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/TewApirat/items-shop-api/pkg/itemManaging/service"
	_itemManagingController "github.com/TewApirat/items-shop-api/pkg/itemManaging/controller"
	_itemShopRepository "github.com/TewApirat/items-shop-api/pkg/itemShop/repository"
)

func (s *echoServer)initItemManagingRouter(){

	router := s.app.Group("v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(
		itemManagingRepository,
		itemShopRepository,
	)
	itemMangingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemMangingController.Creating)
	router.PATCH("/:itemID",itemMangingController.Editing)


}
