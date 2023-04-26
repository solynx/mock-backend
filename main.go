package main

import (
	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/template/html"
	// "github.com/gofiber/fiber/middleware"
	"app/cmd/routers"
	"app/database"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// "fmt"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

func main() {

	app := fiber.New()

	app.Use(
		skip.New(
			cors.New(cors.Config{AllowCredentials: true}),
			func(ctx *fiber.Ctx) bool {
				return strings.Contains(ctx.Get(fiber.HeaderOrigin, ""), ":://solynx.github.io")
			},
		),
	)
	routers.SetupRouters(app)
	database.InitDb()

	app.Listen(":8000")

}
