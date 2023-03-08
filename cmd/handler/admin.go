package handler

import (
	"github.com/gofiber/fiber/v2"
	"app/models/courses"
	// "github.com/google/uuid"
	"app/database"
	// "fmt"
)
func GetAllCourses (c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	
	data := courses.GetAllCourses(db)
	
	// if result {
	// 	return c.JSON(fiber.Map{"status": "success", "message": "Insert data successfully!", "data": nil})
	// }
	return c.JSON(fiber.Map{"error":nil,"message": "Success","status": "false",  "items": data})
}
func CreateCourse (c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want create a course")
}
func ActiveCourse(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want active course")
}
func CourseDetail(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want a course detail" )
}
func UpdateCourse(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want update a course" )
}
func DeleteCourse(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want delete a course" )
}