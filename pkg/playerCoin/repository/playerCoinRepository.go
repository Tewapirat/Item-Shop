package repository

import (
	"github.com/TewApirat/items-shop-api/entities"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
	"gorm.io/gorm"
)
type PlayerCoinRepository interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin)(*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)

}

