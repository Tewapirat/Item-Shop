package repository

import (

	entities "github.com/TewApirat/items-shop-api/entities"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)


type CoinRepositoryMock struct{
	mock.Mock
}

func (m * CoinRepositoryMock)CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin)(*entities.PlayerCoin, error){
	args := m.Called(tx, playerCoinEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *CoinRepositoryMock)Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error){
	args := m.Called(playerID)
	return args.Get(0).(*_playerCoinModel.PlayerCoinShowing), args.Error(1)
}


