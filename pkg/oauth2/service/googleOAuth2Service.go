package service

import (
	"github.com/TewApirat/items-shop-api/entities"
	_adminModel "github.com/TewApirat/items-shop-api/pkg/admin/model"
	_adminRepositoty "github.com/TewApirat/items-shop-api/pkg/admin/repository"
	_playerModel "github.com/TewApirat/items-shop-api/pkg/player/model"
	_playerRepository "github.com/TewApirat/items-shop-api/pkg/player/repository"
)


type googleOAuth2Service struct{
	adminRepository _adminRepositoty.AdminRepository
	playerRepository _playerRepository.PlayerRepository
}

func NewGoogleOAuth2Service (
	adminRepository _adminRepositoty.AdminRepository,
	playerRepositpry _playerRepository.PlayerRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		adminRepository,
		playerRepositpry,
	}
}


func (s *googleOAuth2Service)PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	
	if !s.IsThisGuyIsReallyPlayer(playerCreatingReq.ID) {
	playerEntity := &entities.Player{

		ID: playerCreatingReq.ID,
		Name: playerCreatingReq.Name,
		Email: playerCreatingReq.Email,
		Avatar: playerCreatingReq.Avatar,
	}

	if _, err := s.playerRepository.Creating(playerEntity); err != nil{
		return err
		}
	}

	return nil
}

func (s *googleOAuth2Service)AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {
	if !s.IsThisGuyIsReallyAdmin(adminCreatingReq.ID) {
		adminEntity := &entities.Admin{

			ID:     adminCreatingReq.ID,
			Name:   adminCreatingReq.Name,
			Email:  adminCreatingReq.Email,
			Avatar: adminCreatingReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}
	return player != nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}
	return admin != nil
}

