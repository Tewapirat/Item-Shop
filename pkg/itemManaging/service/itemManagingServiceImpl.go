package service

import (
	"github.com/TewApirat/items-shop-api/entities"
	_itemManagingModel "github.com/TewApirat/items-shop-api/pkg/itemManaging/model"
	_itemManagingRepository "github.com/TewApirat/items-shop-api/pkg/itemManaging/repository"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
)	

type itemManagingServiceImpl struct {

	_itemManagingRepository _itemManagingRepository.ItemManagingRepository

}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository,) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}

func (s *itemManagingServiceImpl)Creating(ItemCreatingReq *_itemManagingModel.ItemCreatingReq) ( *_itemShopModel.Item,error){
	itemEntity := &entities.Item{
		Name:			ItemCreatingReq.Name,
		Description: 	ItemCreatingReq.Description,
		Picture: 		ItemCreatingReq.Picture,
		Price: 			ItemCreatingReq.Price,	
	}

	itemEntityResult, err := s._itemManagingRepository.Creating(itemEntity)
	if err != nil{
		return nil,err
	}
	
	
	
	return itemEntityResult.ToItemModel(),nil
}
