package v1

import (
	"github.com/gofiber/fiber/v2"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
	"time"
)

type ResponseInventory struct {
	ID        int       `json:"id"`
	Store     string    `json:"store"`
	ShoeModel string    `json:"shoe_model"`
	Amount    int       `json:"amount"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseSuggestion struct {
	LowStockStore   string `json:"low_stock_store"`
	HighStockStore  string `json:"high_stock_store"`
	ShoeModel       string `json:"shoe_model"`
	LowStockAmount  int    `json:"low_stock_amount"`
	HighStockAmount int    `json:"high_stock_amount"`
}

func CreateResponseInventory(inventory entity.Inventory) ResponseInventory {
	return ResponseInventory{
		ID:        int(inventory.ID),
		Store:     inventory.Store.Name,
		ShoeModel: inventory.ShoeModel.Name,
		Amount:    inventory.Amount,
		UpdatedAt: inventory.UpdatedAt,
	}
}

func CreateInventorySuggestion(lsInv entity.Inventory, hsInv entity.Inventory) ResponseSuggestion {
	return ResponseSuggestion{
		LowStockStore:   lsInv.Store.Name,
		HighStockStore:  hsInv.Store.Name,
		ShoeModel:       lsInv.ShoeModel.Name,
		LowStockAmount:  lsInv.Amount,
		HighStockAmount: hsInv.Amount,
	}
}

func GetAllInventory(c *fiber.Ctx) error {
	var inventories []entity.Inventory
	sqlite.GetInventories(&inventories)

	responseInventories := []ResponseInventory{}
	for _, inventory := range inventories {
		responseInventory := CreateResponseInventory(inventory)
		responseInventories = append(responseInventories, responseInventory)
	}

	return c.Status(fiber.StatusOK).JSON(responseInventories)
}

func GetInventoryForStore(c *fiber.Ctx) error {
	var inventories []entity.Inventory

	storeId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	sqlite.GetInventoriesByStoreId(&inventories, uint(storeId))

	responseInventories := []ResponseInventory{}
	for _, inventory := range inventories {
		responseInventory := CreateResponseInventory(inventory)
		responseInventories = append(responseInventories, responseInventory)
	}

	return c.Status(fiber.StatusOK).JSON(responseInventories)
}

func InventoryMoveSuggestions(c *fiber.Ctx) error {
	lowStockInventories := []entity.Inventory{}
	sqlite.GetLowStockInventories(&lowStockInventories)

	suggestions := []ResponseSuggestion{}
	for _, lsInv := range lowStockInventories {
		hsInv := entity.Inventory{ShoeModelID: lsInv.ShoeModelID}
		sqlite.GetHighStockInventory(&hsInv)

		suggestion := CreateInventorySuggestion(lsInv, hsInv)
		suggestions = append(suggestions, suggestion)
	}

	return c.Status(fiber.StatusOK).JSON(suggestions)

}
