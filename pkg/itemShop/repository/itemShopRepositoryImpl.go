package repository

import (
	"github.com/TewApirat/items-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/TewApirat/items-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?",false) // Select * from items
	if itemFilter.Name != " " {
		query = query.Where("name like ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != " " {
		query = query.Where("description like ?", "%"+itemFilter.Description+"%")
	}

	offset := int(itemFilter.Page - 1) * int(itemFilter.Size)
	limit := int(itemFilter.Size)




	if err := query.Offset(offset).Limit(limit).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?",false) // Select * from items
	if itemFilter.Name != " " {
		query = query.Where("name like ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != " " {
		query = query.Where("description like ?", "%"+itemFilter.Description+"%")
	}


	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting item failed : %s", err.Error())
		return -1, &_itemShopException.ItemCounting{}
	}

	return count, nil
}
