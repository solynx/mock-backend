package handler

import (
	"github.com/gofiber/fiber/v2"
)
func GetCourceList (c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}