package main

import (
   
    "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/template/html"
    // "github.com/gofiber/fiber/middleware"
    "app/cmd/routers"
)

func main() {
    app := fiber.New()
    
    routers.SetupRouters(app);
    app.Listen(":8000")
}