package repository

import (
	
	"github.com/TewApirat/items-shop-api/entities"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
)
type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin)(*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)

}

