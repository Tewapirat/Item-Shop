package service

import (

	"github.com/TewApirat/items-shop-api/entities"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
	_playerCoinRepositroy "github.com/TewApirat/items-shop-api/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _playerCoinRepositroy.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(
	playerCoinRepository _playerCoinRepositroy.PlayerCoinRepository,
)PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}


func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin,error){
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount: coinAddingReq.Amount,

	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(playerCoinEntity)
	if err != nil {
		return nil, err
		
	}

	playerCoinEntityResult.PlayerID = coinAddingReq.PlayerID

	return  playerCoinEntityResult.ToPlayerCoinModel(),nil

}


func (s *playerCoinServiceImpl)Showing(playerID string) *_playerCoinModel.PlayerCoinShowing{
	
	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin: 0,
		}
	}
	
	return playerCoinShowing
}
