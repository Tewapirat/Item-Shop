package repository

import "github.com/TewApirat/items-shop-api/entities"

type AdminRepository interface{
	Creating(adminEntity *entities.Admin)(*entities.Admin, error) 
	FindByID(admunID string) (*entities.Admin, error) 
	
}