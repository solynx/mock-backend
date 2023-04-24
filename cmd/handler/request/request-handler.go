package request_handler

import (
	"app/database"
	"app/models/menu/requests"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllRequests(c *fiber.Ctx) error {
	db := database.Database
	data := requests.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewRequest(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	request := new(requests.Request)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	db := database.Database
	result := requests.CreateARequest(db, *request)

	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The request created successfully!", "status": true})

	}
	return c.JSON(fiber.Map{"error": nil, "message": "The request created don't successfully!", "status": false})
}
func UpdateRequest(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	request := new(requests.Request)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	request.UpdatedAt = time.Now()
	db := database.Database
	result := requests.UpdateRequestById(db, *request)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})

}
func RemoveRequest(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	request := new(requests.Request)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.Database

	result := requests.RemoveRequest(db, *request)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data deleted don't successfully!", "status": false})
}
