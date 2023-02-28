package routers

import (
	// "app/cmd/controllers"
	"github.com/gofiber/fiber/v2"
	"app/cmd/handler"
	//  "app/cmd/middlewares"
	
	 "github.com/gofiber/fiber/v2/middleware/logger"
)
func SetupRouters(app *fiber.App){
	admin := app.Group("/admin",logger.New())
	admin.Get("/cource.json",handler.GetCourceList)
	
}