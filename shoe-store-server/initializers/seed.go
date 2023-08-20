package initializers

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
	"strconv"
)

func SeedDB() {
	log.Info("Seeding DB...")

	loadStores()
	loadShoeModelsAndInventory()

	log.Info("Seeding complete!")
}

func loadStores() {
	for _, storeName := range getShoeStores() {
		sqlite.CreateStore(&entity.Store{Name: storeName})
	}
}

func loadShoeModelsAndInventory() {
	var stores []entity.Store
	sqlite.GetStores(&stores)

	for _, smName := range getShoeModels() {
		shoe := entity.ShoeModel{Name: smName}
		sqlite.CreateShoeModel(&shoe)

		for _, store := range stores {
			amt, _ := strconv.Atoi(os.Getenv("MAX_INVENTORY"))
			inventory := entity.Inventory{
				Amount:    amt,
				ShoeModel: shoe,
				Store:     store,
			}
			sqlite.CreateInventory(&inventory)
		}
	}
}

func getShoeStores() []string {
	return []string{
		"ALDO Centre Eaton",
		"ALDO Destiny USA Mall",
		"ALDO Pheasant Lane Mall",
		"ALDO Holyoke Mall",
		"ALDO Maine Mall",
		"ALDO Crossgates Mall",
		"ALDO Burlington Mall",
		"ALDO Solomon Pond Mall",
		"ALDO Auburn Mall",
		"ALDO Waterloo Premium Outlets",
	}
}

func getShoeModels() []string {
	return []string{
		"ADERI",
		"MIRIRA",
		"CAELAN",
		"BUTAUD",
		"SCHOOLER",
		"SODANO",
		"MCTYRE",
		"CADAUDIA",
		"RASIEN",
		"WUMA",
		"GRELIDIEN",
		"CADEVEN",
		"SEVIDE",
		"ELOILLAN",
		"BEODA",
		"VENDOGNUS",
		"ABOEN",
		"ALALIWEN",
		"GREG",
		"BOZZA",
	}
}
