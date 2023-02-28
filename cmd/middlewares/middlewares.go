package middlewares

import (
    "github.com/gofiber/fiber/v2"
	"fmt"
)

func LoggerMiddleware(c *fiber.Ctx) error {
    fmt.Println("Request received")
    return c.Next()
}