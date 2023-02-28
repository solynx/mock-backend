package routers

import (
	"github.com/gofiber/fiber/v2"
	 "app/cmd/middlewares"
	 "app/cmd/controllers"
)
func SetupRouters(app *fiber.App){
	app.Get("/",middlewares.LoggerMiddleware,mainController.HomepageController)
}