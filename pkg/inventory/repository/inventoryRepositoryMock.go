package repository

import (
	entities "github.com/TewApirat/items-shop-api/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type InventoryRepositoryMock struct{
	mock.Mock
}

func (m *InventoryRepositoryMock)Filling(tx *gorm.DB, playerID string, itemID uint64, qty int)([]*entities.Inventory, error){
	args := m.Called(tx, playerID, itemID, qty)
	return args.Get(0).([]*entities.Inventory),args.Error(1)
}

func (m *InventoryRepositoryMock)Removing(tx *gorm.DB, playerID string, itemID uint64, limit int)error{
	args := m.Called(tx, playerID, itemID)
	return args.Error(0)
}

func (m *InventoryRepositoryMock)PlayerItemCounting(playerId string, itemID uint64) int64{
	args := m.Called(playerId, itemID)
	return args.Get(0).(int64)
}

func (m *InventoryRepositoryMock)Listing(playerId string)([]*entities.Inventory, error){
	args := m.Called(playerId)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}