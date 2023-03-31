package collection_handler

import (
	"app/database"
	"app/models/menu/collections"
	"app/models/menu/responses"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllCollections(c *fiber.Ctx) error {
	db := database.Database
	data := collections.GetMenu(db)
	c.Append("Access-Control-Allow-Origin", "*")
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func GetACollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(collections.Collection)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	result := collections.GetCollection(db, *collection)
	return c.JSON(fiber.Map{"error": nil, "data": result, "status": true})
}
func FormatCollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(responses.ResponseFormat)

	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db.Table("response_formats").Debug().Find(&collection)
	sDec, _ := base64.StdEncoding.DecodeString(collection.OriginalRequest)

	fmt.Println(string(sDec))
	// Giải mã mảng byte thành map[string][]responses.Response

	return c.JSON(fiber.Map{"error": nil, "data": collection, "status": true})
}
func CreateNewCollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(collections.Collection)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	collection.CreatedAt = time.Now()
	collection.UpdatedAt = time.Now()
	collection.UserId = 1

	result := collections.CreateACollection(db, *collection)

	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The collection created successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The collection created don't successfully!", "status": false})

}
func UpdateCollection(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	collection := new(collections.Collection)
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	collection.UpdatedAt = time.Now()
	db := database.Database
	result := collections.UpdateCollectionById(db, *collection)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})

}
func RemoveCollection(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	collection := new(collections.Collection)

	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.Database
	data := collections.Collection{}
	data.ID = collection.ID
	result := collections.RemoveCollection(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
