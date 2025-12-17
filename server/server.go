package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/TewApirat/items-shop-api/config"
	"github.com/TewApirat/items-shop-api/databases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_adminRepository "github.com/TewApirat/items-shop-api/pkg/admin/repository"
	_playerRepository "github.com/TewApirat/items-shop-api/pkg/player/repository"
	_oauth2Controller "github.com/TewApirat/items-shop-api/pkg/oauth2/controller"
	_oauth2Service "github.com/TewApirat/items-shop-api/pkg/oauth2/service"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}

	})

	return server

}
func (s *echoServer) Start() {

	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeOutMiddleware)

	authorizingMiddleware := s.getAuthorzingMiddleware()



	s.app.GET("/v1/health", s.healthCheck)
	

	s.initOAuth2Router()
	s.initItemShopRouter(authorizingMiddleware)
	s.initItemManagingRouter(authorizingMiddleware)
	s.initPlayerCoinRouter(authorizingMiddleware)
	s.initinventoryRouter(authorizingMiddleware)


	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening()

}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("ERROR %s", err.Error())
	}
}


func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal){

	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("Shutting down server...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal("Error: %s", err.Error())
	}

}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")

}



func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return  middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout: timeout *time.Second,
	})
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleware(bodylimit string) echo.MiddlewareFunc{
	return middleware.BodyLimit(bodylimit)
}

func (s * echoServer)getAuthorzingMiddleware() *authorizingMiddleware{
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(adminRepository,playerRepository)
	oauth2Controller := _oauth2Controller.NewgoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		s.app.Logger,
	)

	return &authorizingMiddleware{
		oauth2Controller: 	oauth2Controller,
		oauth2Config: 		s.conf.OAuth2,
		logger:				s.app.Logger,
	}

}