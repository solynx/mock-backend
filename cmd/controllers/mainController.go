package mainController

import (
	"github.com/gofiber/fiber/v2")

func HomepageController ( c *fiber.Ctx) error {
	return c.SendString("Hello i'm controller and router 1 middle2 controller 3")
} 