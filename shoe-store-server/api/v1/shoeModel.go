package v1

import (
	"github.com/gofiber/fiber/v2"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
	"sort"
)

type ResponseShoeModel struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	TotalSales int    `json:"total_sales"`
}

func CreateResponseShoeModel(shoeModel entity.ShoeModel) ResponseShoeModel {
	if shoeModel.ID == 0 || shoeModel.Name == "" {
		return ResponseShoeModel{}
	}

	totalSales := getTotalSalesByShoeModel(shoeModel.ID)
	return ResponseShoeModel{
		ID:         shoeModel.ID,
		Name:       shoeModel.Name,
		TotalSales: totalSales,
	}
}

func GetShoeModels(c *fiber.Ctx) error {
	var shoeModels []entity.ShoeModel
	sqlite.GetShoeModels(&shoeModels)

	responseShoeModels := []ResponseShoeModel{}
	for _, shoeModel := range shoeModels {
		responseShoeModel := CreateResponseShoeModel(shoeModel)
		responseShoeModels = append(responseShoeModels, responseShoeModel)
	}

	sort.Slice(responseShoeModels, func(i, j int) bool {
		return responseShoeModels[i].TotalSales >
			responseShoeModels[j].TotalSales
	})
	return c.Status(fiber.StatusOK).JSON(responseShoeModels)
}

func GetShoeModel(c *fiber.Ctx) error {
	var shoeModel entity.ShoeModel

	id, paramErr := c.ParamsInt("id")
	if paramErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	shoeModel.ID = uint(id)

	if err := sqlite.GetShoeModel(&shoeModel); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "shoe model not found"})
	}

	responseShoeModel := CreateResponseShoeModel(shoeModel)
	return c.Status(fiber.StatusOK).JSON(responseShoeModel)
}

func CreateShoeModel(c *fiber.Ctx) error {
	var shoeModel entity.ShoeModel

	if err := c.BodyParser(&shoeModel); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	sqlite.CreateShoeModel(&shoeModel)
	responseShoeModel := CreateResponseShoeModel(shoeModel)
	return c.Status(fiber.StatusCreated).JSON(responseShoeModel)
}

func getTotalSalesByShoeModel(shoeModelId uint) int {
	var totalSales int64
	sqlite.GetNumberOfSalesByShoeModelId(&totalSales, shoeModelId)
	return int(totalSales)
}
