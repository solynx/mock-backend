package request_handler

import (
	"app/database"
	"app/models/menu/requests"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllRequests(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := requests.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewRequest(c *fiber.Ctx) error {
	request := new(requests.Request)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	request.ID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	request.CollectionId = "369e28c3-bf4c-4066-84e9-1180a3fd285f"
	db := database.ConnectDBLocal()
	data := requests.CreateARequest(db, *request)
	_, err := uuid.Parse(data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The request created don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The request created successfully!", "status": true, "item": data})
}
func UpdateRequest(c *fiber.Ctx) error {
	request := new(requests.Request)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	request.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := requests.UpdateRequestById(db, *request)
	if result {
		return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
}
func RemoveRequest(c *fiber.Ctx) error {
	id := c.Query("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": false, "message": "id " + id + " must be uuid type and exists"})
	}
	db := database.ConnectDBLocal()
	data := requests.Request{}
	data.ID = uuid
	result := requests.RemoveRequest(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
