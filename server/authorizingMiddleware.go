package server

import (
	"github.com/TewApirat/items-shop-api/config"
	_oauth2Controller "github.com/TewApirat/items-shop-api/pkg/oauth2/controller"
	"github.com/labstack/echo/v4"
	
)

type authorizingMiddleware struct {
	oauth2Controller _oauth2Controller.OAuth2Controller
	oauth2Config 	*config.OAuth2
	logger			echo.Logger
}


func (m * authorizingMiddleware)PlayerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return  func(pctx echo.Context) error {
		return m.oauth2Controller.PlayerAuthorizing(pctx, next)
	}
}

func (m * authorizingMiddleware)AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return  func(pctx echo.Context) error {
		return m.oauth2Controller.AdminAuthorizing(pctx, next)
	}
}

