package service

import (
	"github.com/TewApirat/items-shop-api/entities"
	_inventoryRepository "github.com/TewApirat/items-shop-api/pkg/inventory/repository"
	_itemShopException "github.com/TewApirat/items-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/TewApirat/items-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/TewApirat/items-shop-api/pkg/itemShop/repository"
	_playerCoinModel "github.com/TewApirat/items-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/TewApirat/items-shop-api/pkg/playerCoin/repository"
	"github.com/labstack/echo/v4"
)

type itemShopServiceImpl struct {

	itemShopRepository _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger

	
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
	inventoryRepository  _inventoryRepository.InventoryRepository,
	logger               echo.Logger,
	) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		logger,
	}
}

func (s * itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error){
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil{
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page
	totalPage := s.totalPageCalculation(itemCounting, size)

	return s.toItemResultResponse(itemList, page, totalPage), nil
}
// 1. Find Item by ID
// 2. Total Price calculation
// 3. Check player Coin
// 4. Purchase History Recording
// 5. Coin Reducting
// 6. Inventory Filling
// 7. Return Player Coin
func (s *itemShopServiceImpl)Buying(buyingReq *_itemShopModel.BuyingReq)(*_playerCoinModel.PlayerCoin, error){
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil{
		return nil,err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(),buyingReq.Quantity)
	
	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil{
		return nil, err
	}


	tx := s.itemShopRepository.TransactionBegin()

	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID: buyingReq.PlayerID,
		ItemID: buyingReq.ItemID,
		ItemName: itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice: itemEntity.Price,
		ItemPicture: itemEntity.Picture,
		Quantity: buyingReq.Quantity,
	})

	if err != nil{
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Info("Purchase history recorded: %s",purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount: -totalPrice,
	})
	if err != nil{
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("Player coin reducted: %d",playerCoin.Amount)

	inventoryEntity, err := s.inventoryRepository.Filling(tx,
		buyingReq.PlayerID,
		buyingReq.ItemID,
		int(buyingReq.Quantity),
	)
	if err != nil{
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Inventory filled: %d", len(inventoryEntity))

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil{
		return nil,err
	}

	return playerCoin.ToPlayerCoinModel(),nil
}


// 1. Find Item by ID
// 2. Total Price calculation
// 3. Check player Item
// 4. Purchase History Recording
// 5. Coin Adding
// 6. Inventory Removing
// 7. Return Player Coin

func (s *itemShopServiceImpl)Selling(sellinggReq *_itemShopModel.BuyingReq)(*_playerCoinModel.PlayerCoin, error){
	return nil,nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}
	return int64(totalPage)
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: 		itemModelList,
		Paginate:		_itemShopModel.PaginateResult{
			Page: 		page,
			TotalPage: 	totalPage,
		},
	}
}

func (s * itemShopServiceImpl) totalPriceCalculation(item * _itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil{
		return err
	}

	if playerCoin.Coin < totalPrice{
		s.logger.Error("Player coin is not enough")
		return &_itemShopException.CoinNotEnough{}
	}
	
	return nil
}