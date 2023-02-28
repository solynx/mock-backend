package handler

import (
	"github.com/gofiber/fiber/v2"
)
func GetCourceList (c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
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