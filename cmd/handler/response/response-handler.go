package response_handler

import (
	"app/database"
	"app/models/menu/responses"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllResponses(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := responses.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewResponse(c *fiber.Ctx) error {
	response := new(responses.Response)
	if err := c.BodyParser(&response); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	response.ID = uuid.New()
	response.CreatedAt = time.Now()
	response.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	data := responses.CreateAResponse(db, *response)
	_, err := uuid.Parse(data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The response created don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The response created successfully!", "status": true, "item": data})
}
func Updateresponse(c *fiber.Ctx) error {
	response := new(responses.Response)
	if err := c.BodyParser(&response); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	response.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := responses.UpdateResponseById(db, *response)
	if result {
		return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
}
func RemoveResponse(c *fiber.Ctx) error {
	id := c.Query("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": false, "message": "id " + id + " must be uuid type and exists"})
	}
	db := database.ConnectDBLocal()
	data := responses.Response{}
	data.ID = uuid
	result := responses.RemoveResponse(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
