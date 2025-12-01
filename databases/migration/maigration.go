package main

import (
	// "fmt"

	"github.com/TewApirat/items-shop-api/config"
	"github.com/TewApirat/items-shop-api/databases"
	"github.com/TewApirat/items-shop-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGettings()
	db := databases.NewPostgresDatabase(conf.Database)


	tx := db.ConnectionGetting().Begin()

	// fmt.Println(db.ConnectionGetting())

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseHistoryMigration(tx)

	tx.Commit()
	if tx.Error != nil{
		tx.Rollback()
		panic(tx.Error)
	}
}

func playerMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB){
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}

