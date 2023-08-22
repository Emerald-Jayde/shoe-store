package v1

import (
	"github.com/gofiber/fiber/v2"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
	"sort"
)

type ResponseStore struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	TotalSales  int    `json:"total_sales"`
	BestSeller  string `json:"best_seller"`
	WorstSeller string `json:"worst_seller"`
}

func CreateResponseStore(store entity.Store) ResponseStore {
	if store.ID == 0 || store.Name == "" {
		return ResponseStore{}
	}

	totalSales := getTotalSalesByStore(store.ID)
	s := sqlite.GetBestAndWorstShoeModelSalesByStoreId(store.ID)
	bestSeller := sqlite.GetShoeModelNameById(uint(s["best"][0]))
	worstSeller := sqlite.GetShoeModelNameById(uint(s["worst"][0]))

	return ResponseStore{
		ID:          store.ID,
		Name:        store.Name,
		TotalSales:  totalSales,
		BestSeller:  bestSeller,
		WorstSeller: worstSeller,
	}
}

func GetStores(c *fiber.Ctx) error {
	var stores []entity.Store
	sqlite.GetStores(&stores)

	responseStores := []ResponseStore{}
	for _, store := range stores {
		responseStore := CreateResponseStore(store)
		responseStores = append(responseStores, responseStore)
	}

	sort.Slice(responseStores, func(i, j int) bool {
		return responseStores[i].TotalSales >
			responseStores[j].TotalSales
	})
	return c.Status(fiber.StatusOK).JSON(responseStores)
}

func GetStore(c *fiber.Ctx) error {
	var store entity.Store

	id, paramErr := c.ParamsInt("id")
	if paramErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	store.ID = uint(id)
	if err := sqlite.GetStore(&store); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "store not found"})
	}

	responseStore := CreateResponseStore(store)
	return c.Status(fiber.StatusOK).JSON(responseStore)
}

func CreateStore(c *fiber.Ctx) error {
	var store entity.Store

	if err := c.BodyParser(&store); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	sqlite.CreateStore(&store)
	responseStore := CreateResponseStore(store)
	return c.Status(fiber.StatusCreated).JSON(responseStore)
}

func getTotalSalesByStore(shoeModelId uint) int {
	var totalSales int64
	sqlite.GetNumberOfSalesByStoreId(&totalSales, shoeModelId)
	return int(totalSales)
}
