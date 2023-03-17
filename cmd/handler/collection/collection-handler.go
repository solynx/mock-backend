package collection_handler

import (
	"app/database"
	"app/models/menu/collections"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllCollections(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := collections.GetMenu(db)
	c.Append("Access-Control-Allow-Origin", "*")
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewCollection(c *fiber.Ctx) error {
	collection := new(collections.Collection)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	fmt.Printf("%+v\n", collection)
	collection.CreatedAt = time.Now()
	collection.UpdatedAt = time.Now()
	collection.UserId = 1
	db := database.ConnectDBLocal()
	data := collections.CreateACollection(db, *collection)
	_, err := uuid.Parse(data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The collection created don't successfully!", "status": false})
	}

	return c.JSON(fiber.Map{"error": nil, "message": "The collection created successfully!", "status": true, "item": data})
}
func UpdateCollection(c *fiber.Ctx) error {
	collection := new(collections.Collection)
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	collection.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := collections.UpdateCollectionById(db, *collection)
	if result {
		return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
}
func RemoveCollection(c *fiber.Ctx) error {
	id := c.Query("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": false, "message": "id " + id + " must be uuid type and exists"})
	}
	db := database.ConnectDBLocal()
	data := collections.Collection{}
	data.ID = uuid
	result := collections.RemoveCollection(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}