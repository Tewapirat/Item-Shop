package repository

import "github.com/TewApirat/items-shop-api/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item)(*entities.Item , error)

}
