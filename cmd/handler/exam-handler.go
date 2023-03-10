package handler

import (
	"app/database"
	"app/models/courses"
	"app/models/exams"
	"fmt"
	"strconv"
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
	/*
			 {
		        "Id": "0a281988-a933-48b9-9e49-e0a9be42f841",
		        "name": "Bulma",
		        "type": "online",
		        "category_id": 1,
		        "content": "This is content 123",
		        "status": "this is status1",
		        "UserId": "a61a0fbb-7aee-43c7-b7a7-df9cf45a30f6",
		        "question_id": 1,
		        "CreatedAt": "2023-03-10T10:57:33.574364369+07:00",
		        "UpdatedAt": "2023-03-10T10:57:33.574364458+07:00"
		    }
	*/
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
func CourseDetail(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	db := database.ConnectDBLocal()
	data, rows := courses.GetCourseById(db, id)
	if rows > 0 {
		return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "the data created NOT successfully!", "data": data})
	}
	return c.JSON(fiber.Map{"error": "not found", "status": "failed", "message": "The course id " + strconv.Itoa(int(id)) + " not found !"})
}
func UpdateCourse(c *fiber.Ctx) error {
	type NewC struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Price       string `json:"price"`
		Description *int   `json:"description", omitempty`
	}

	var newCourse NewC
	test := NewC{
		ID:    1,
		Title: "hello",
		Price: "300",
	}
	fmt.Println(test)
	if err := c.BodyParser(&newCourse); err != nil {
		fmt.Println("=========>", err)
		return c.JSON(fiber.Map{"error": "not found", "status": false, "message": "false"})
	}
	fmt.Println(newCourse)

	// 	return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "the data created successfully!"})

	// db := database.ConnectDBLocal()
	// result := courses.UpdateCourse(db, payload )
	// if result {
	// return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "the data created successfully!"})
	// }
	return c.JSON(fiber.Map{"error": nil, "status": "failed", "message": "the data created NOT successfully!", "data": nil})

}
func DeleteCourse(c *fiber.Ctx) error {
	//need softdel
	id := uint32(c.QueryInt("id"))
	course := courses.Course{Id: id}
	db := database.ConnectDBLocal()
	err, rowsNum := courses.DeleteCourse(db, &course)
	if rowsNum > 0 {
		return c.JSON(fiber.Map{"error": nil, "status": "success", "message": "The course id " + strconv.Itoa(int(id)) + " deleted successfully!", "data": nil})
	}

	return c.JSON(fiber.Map{"error": err, "status": "failed", "message": "The course id " + strconv.Itoa(int(id)) + " not found !"})
}
