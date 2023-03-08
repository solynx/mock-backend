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
)

func main() {
    
    app := fiber.New()
        routers.SetupRouters(app);
        database.InitDb()
    app.Listen(":8000")
}