package response_handler

import (
	"app/database"
	"app/models/menu/responses"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllResponses(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := responses.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewResponse(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	response := new(responses.Response)
	if err := c.BodyParser(&response); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	response.CreatedAt = time.Now()
	response.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := responses.CreateAResponse(db, *response)

	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The response created successfully!", "status": true})
	}

	return c.JSON(fiber.Map{"error": nil, "message": "The response created don't successfully!", "status": false})
}
func Updateresponse(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	response := new(responses.Response)
	if err := c.BodyParser(&response); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	response.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := responses.UpdateResponseById(db, *response)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
	}

	return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})
}
func RemoveResponse(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	response := new(responses.Response)
	if err := c.BodyParser(&response); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.ConnectDBLocal()

	result := responses.RemoveResponse(db, *response)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
