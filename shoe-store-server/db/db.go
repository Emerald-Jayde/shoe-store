package db

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shoe-store-server/entity"
	"shoe-store-server/helpers"
)

type DatabaseInstance struct {
	Db *gorm.DB
}

var Database DatabaseInstance

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("shoe-store.db"), &gorm.Config{})
	helpers.HandleError("Failed to connect to SQLite database", err, true)

	log.Info("Connected to the database successfully!")

	log.Info("Running migrations...")
	err = db.AutoMigrate(&entity.Store{}, &entity.ShoeModel{}, &entity.Inventory{}, &entity.Sale{})
	helpers.HandleError("Migration failed", err, true)

	log.Info("DB Migrations complete!")

	Database = DatabaseInstance{Db: db}
}

func GetDBInstance() *gorm.DB {
	return Database.Db
}
