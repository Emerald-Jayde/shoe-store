package initializers

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shoe-store-server/entity"
	"shoe-store-server/helpers"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var Database DatabaseInstance

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("shoeStore.db"), &gorm.Config{})
	helpers.HandleError("Failed to connect to SQLite database", err, true)

	log.Info("Connected to the database successfully!")

	log.Info("Running migrations...")
	db.AutoMigrate(&entity.Store{}, &entity.ShoeModel{}, &entity.Inventory{})
	log.Info("DB Migrations complete!")

	Database = DatabaseInstance{DB: db}
}
