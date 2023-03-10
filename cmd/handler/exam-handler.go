package handler

import (
	"app/database"
	"app/models/exams"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllExam(c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data := exams.GetExams(db)
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": "false", "items": data})
}
func CreateExam(c *fiber.Ctx) error {
	exam := new(exams.CreateExamStruct)
	if err := c.BodyParser(&exam); err != nil {
		return err
	}
	exam.Id = uuid.New()
	exam.UserId = uuid.New()
	exam.CreatedAt = time.Now()
	exam.UpdatedAt = time.Now()

	db := database.ConnectDBLocal()

	result := exams.CreateExam(db, *exam)
	if result != nil {
		return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "the data created NOT successfully!"})
	}
	return c.JSON(fiber.Map{"error": nil, "status": "success", "message": "the data created successfully!", "data": exam})
}
func ActiveCourse(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want active course")
}
func GetExamById(c *fiber.Ctx) error {
	id := c.Query("id")
	db := database.ConnectDBLocal()
	data, rows := exams.GetExamById(db, id)
	if rows > 0 {
		return c.JSON(fiber.Map{"error": nil, "status": "success", "message": "Get exam id " + id + " successfully!", "data": data})
	}
	return c.JSON(fiber.Map{"error": "not found", "status": "failed", "message": "The exam id " + id + " not found !"})
}
func UpdateExam(c *fiber.Ctx) error {
	id := c.Query("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "id " + id + " must be uuid type and exits", "data": nil})
	}
	exam := new(exams.UpdateExamStruct)
	if err := c.BodyParser(&exam); err != nil {
		return err
	}
	exam.Id = id
	exam.UpdatedAt = time.Now()

	db := database.ConnectDBLocal()
	result := exams.UpdateExam(db, *exam)
	if result {
		return c.JSON(fiber.Map{"error": nil, "status": "success", "message": "the data updated successfully!", "data": exam})
	}
	return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "the data updated failed!", "data": nil})

}
func DeleteExam(c *fiber.Ctx) error {
	id := c.Query("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "id " + id + " must be uuid type and exits", "data": nil})
	}
	data := exams.DeleteExamStruct{}
	data.Id = id
	db := database.ConnectDBLocal()
	result := exams.DeleteExam(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "status": "success", "message": "The exam id " + id + " deleted successfully!", "data": nil})
	}

	return c.JSON(fiber.Map{"error": err, "status": "failed", "message": "The exam id " + id + " not found !"})
}
