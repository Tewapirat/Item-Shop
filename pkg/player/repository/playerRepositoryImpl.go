package repository

import (
	"github.com/TewApirat/items-shop-api/databases"
	"github.com/TewApirat/items-shop-api/entities"
	"github.com/labstack/echo/v4"
	
	_playerException "github.com/TewApirat/items-shop-api/pkg/player/exception"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}


func NewPlayerRepositoryImpl(
	db databases.Database,
	Logger echo.Logger,
) PlayerRepository {
	return &playerRepositoryImpl{
		db:     db,
		logger: Logger,
	}
}

func (r *playerRepositoryImpl)Creating(playerEntity *entities.Player)(*entities.Player, error){
	player := new(entities.Player)
	
	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Error("createing player failed : %s",err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}
	return player, nil
}

func (r *playerRepositoryImpl)FindByID(playerID string)(*entities.Player, error){
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?",playerID).First(player).Error; err != nil {

		r.logger.Errorf("Find player by ID failed %s",err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}

	}
	return player, nil
}
