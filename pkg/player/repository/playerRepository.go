package repository

import "github.com/TewApirat/items-shop-api/entities"

type PlayerRepository interface{
	Creating(playEntity *entities.Player)(*entities.Player, error) 
	FindByID(playerID string) (*entities.Player, error) 

}