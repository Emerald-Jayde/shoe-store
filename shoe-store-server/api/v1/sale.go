package v1

import (
	"github.com/gofiber/fiber/v2"
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
	storeName := sqlite.GetStoreNameById(sale.StoreID)
	shoeModelName := sqlite.GetShoeModelNameById(sale.ShoeModelID)

	if storeName == "" || shoeModelName == "" {
		return ResponseSale{}
	}

	return ResponseSale{
		Store:     storeName,
		ShoeModel: shoeModelName,
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

func GetLastXSales(c *fiber.Ctx) error {
	var sales []entity.Sale

	limit, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	sqlite.GetLatestSales(&sales, limit)

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
