package routers

import (
	// "app/cmd/controllers"
	"github.com/gofiber/fiber/v2"
	"app/cmd/handler"

	
	 "github.com/gofiber/fiber/v2/middleware/logger"
	
)
func SetupRouters(app *fiber.App){
	
	admin := app.Group("/admin",logger.New())
	// get all course
	admin.Get("/course.json",handler.GetAllCourses)
	
	//active course
	admin.Post("/course/actived.json",handler.ActiveCourse)
	//C - create
	admin.Post("/course.json",handler.CreateCourse)
	//R - read
	admin.Get("/course/details.json",handler.CourseDetail)
	//U - update
	admin.Patch("course.json",handler.UpdateCourse)
	//D - delete
	admin.Delete("course.json",handler.DeleteCourse)
}