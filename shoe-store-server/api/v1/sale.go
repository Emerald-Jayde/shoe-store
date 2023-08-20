package v1

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
	"time"
)

type ResponseSale struct {
	Store     string    `json:"store"`
	ShoeModel string    `json:"shoe_model"`
	NewAmount int       `json:"new_amount"`
	OldAmount int       `json:"old_amount"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseSale(sale entity.Sale) ResponseSale {
	store := entity.Store{
		Model: gorm.Model{ID: sale.StoreID},
	}
	if err := sqlite.GetStoreName(&store); err != nil {
		return ResponseSale{}
	}

	shoeModel := entity.ShoeModel{
		Model: gorm.Model{ID: sale.ShoeModelID},
	}
	if err := sqlite.GetShoeModelName(&shoeModel); err != nil {
		return ResponseSale{}
	}

	return ResponseSale{
		Store:     store.Name,
		ShoeModel: shoeModel.Name,
		NewAmount: sale.NewInventory,
		OldAmount: sale.OldInventory,
		CreatedAt: sale.CreatedAt,
	}
}

func GetSales(c *fiber.Ctx) error {
	var sales []entity.Sale
	sqlite.GetSales(&sales)

	responseSales := []ResponseSale{}
	for _, sale := range sales {
		responseSale := CreateResponseSale(sale)
		responseSales = append(responseSales, responseSale)
	}

	return c.Status(fiber.StatusOK).JSON(responseSales)
}

func GetSalesForStore(c *fiber.Ctx) error {
	var sales []entity.Sale

	storeId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	sqlite.GetSalesByStoreId(&sales, uint(storeId))

	responseSales := []ResponseSale{}
	for _, sale := range sales {
		responseSale := CreateResponseSale(sale)
		responseSales = append(responseSales, responseSale)
	}

	return c.Status(fiber.StatusOK).JSON(responseSales)
}
