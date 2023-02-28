package main

import (
   
    "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/template/html"
   
    "app/cmd/routers"
)

func main() {
    app := fiber.New()
    routers.SetupRouters(app);
    app.Listen(":3000")
}