package server

import (
	_itemManagingRepository "github.com/TewApirat/items-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/TewApirat/items-shop-api/pkg/itemManaging/service"
	_itemManagingController "github.com/TewApirat/items-shop-api/pkg/itemManaging/controller"
)

func (s *echoServer)initItemManagingRouter(){

	router := s.app.Group("v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemMangingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemMangingController.Creating)


}
