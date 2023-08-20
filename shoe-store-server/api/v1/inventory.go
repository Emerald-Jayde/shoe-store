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

func CreateResponseInventory(inventory entity.Inventory) ResponseInventory {
	return ResponseInventory{
		ID:        int(inventory.ID),
		Store:     inventory.Store.Name,
		ShoeModel: inventory.ShoeModel.Name,
		Amount:    inventory.Amount,
		UpdatedAt: inventory.UpdatedAt,
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
