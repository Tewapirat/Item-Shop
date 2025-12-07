package main

import (
	"github.com/TewApirat/items-shop-api/config"
	"github.com/TewApirat/items-shop-api/databases"
	"github.com/TewApirat/items-shop-api/server"
)

func main() {
	conf := config.ConfigGettings()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
