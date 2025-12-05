package repository

import (
	"github.com/TewApirat/items-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_itemManagingException "github.com/TewApirat/items-shop-api/pkg/itemManaging/exception"
)

type itemManagingRepositoryImpl struct{
	db	*gorm.DB
	logger echo.Logger

}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImpl)Creating(itemEntity *entities.Item)(*entities.Item,error){
	item := new(entities.Item)
	
	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Error("createing item failed : %s",err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}
	return item,nil
}