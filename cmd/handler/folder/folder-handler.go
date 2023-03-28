package folder_handler

import (
	"app/database"
	"app/models/menu/folders"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllFolders(c *fiber.Ctx) error {
	db := database.Database
	data := folders.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewFolder(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")

	folder := new(folders.Folder)
	if err := c.BodyParser(&folder); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The json data not valid", "status": false})
	}
	folder.CreatedAt = time.Now()
	folder.UpdatedAt = time.Now()

	db := database.Database
	data := folders.CreateAFolder(db, *folder)
	_, err := uuid.Parse(data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The folder created don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The folder created successfully!", "status": true})
}
func UpdateFolder(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	folder := new(folders.Folder)
	if err := c.BodyParser(&folder); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	folder.UpdatedAt = time.Now()
	db := database.Database
	result := folders.UpdateFolderById(db, *folder)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})

}
func RemoveFolder(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")

	folder := new(folders.Folder)
	if err := c.BodyParser(&folder); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.Database

	result := folders.RemoveFolder(db, *folder)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
