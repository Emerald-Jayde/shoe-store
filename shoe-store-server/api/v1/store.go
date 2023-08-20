package v1

import (
	"github.com/gofiber/fiber/v2"
	"shoe-store-server/entity"
	"shoe-store-server/repository/sqlite"
)

type ResponseStore struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateResponseStore(store entity.Store) ResponseStore {
	if store.ID == 0 || store.Name == "" {
		return ResponseStore{}
	}

	return ResponseStore{
		ID:   store.ID,
		Name: store.Name,
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

	return c.Status(fiber.StatusOK).JSON(responseStores)
}

// TODO: for repository decoupling
// 1. dont call db directly, call repository.sqlite.store.get
// 1.1 return a real entity
// 2. remove gorm from entity, create the gorm entities in a separate folder
// 2.1. the gorm entity will have a constructor to make a GO entity
// 3. in the api calls, for CreateResponse, you use the real entity
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
