package folder_handler

import (
	"app/database"
	"app/models/menu/folders"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllFolders(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := folders.GetMenu(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func CreateNewFolder(c *fiber.Ctx) error {
	folder := new(folders.Folder)
	if err := c.BodyParser(&folder); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The json data not valid", "status": false})
	}
	folder.ID = uuid.New()
	folder.CreatedAt = time.Now()
	folder.UpdatedAt = time.Now()

	db := database.ConnectDBLocal()
	data := folders.CreateAFolder(db, *folder)
	_, err := uuid.Parse(data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The folder created don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The folder created successfully!", "status": true, "item": data})
}
func UpdateFolder(c *fiber.Ctx) error {
	folder := new(folders.Folder)
	if err := c.BodyParser(&folder); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	folder.UpdatedAt = time.Now()
	db := database.ConnectDBLocal()
	result := folders.UpdateFolderById(db, *folder)
	if result {
		return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
}
func RemoveFolder(c *fiber.Ctx) error {
	id := c.Query("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": false, "message": "id " + id + " must be uuid type and exists"})
	}
	db := database.ConnectDBLocal()
	data := folders.Folder{}
	data.ID = uuid
	result := folders.RemoveFolder(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
