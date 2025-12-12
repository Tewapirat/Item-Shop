package service


import (
_playerModel "github.com/TewApirat/items-shop-api/pkg/player/model"
_adminModel "github.com/TewApirat/items-shop-api/pkg/admin/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}