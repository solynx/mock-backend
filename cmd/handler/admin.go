package handler

import (
	"github.com/gofiber/fiber/v2"
	"app/models/courses"
	"strconv"
	"app/database"
	"fmt"
	"github.com/gosimple/slug"
)
func GetAllCourses (c *fiber.Ctx) error {
	db := database.ConnectDBLocal()
	data,err := courses.GetAllCourses(db)
	
	if err != nil {
		return c.JSON(fiber.Map{"error":err,"status": "success", "message": "failed", "data": nil})
	}
	return c.JSON(fiber.Map{"error":nil,"message": "Success","status": "false",  "items": data})
}
func CreateCourse (c *fiber.Ctx) error {
	
	course := courses.Course{
							Title:c.Query("title"),
							Description : c.Query("description") ,
							Price : float32(c.QueryInt("price")),
							SalePrice: float32(c.QueryInt("sale_price")),
							ImageUrl:c.Query("image_url","empty"),
							Slug: slug.Make(c.Query("title")),
							OwnerId: uint32(c.QueryInt("owner_id",1)),}
	db := database.ConnectDBLocal()
	fmt.Println(course)

	result := courses.CreateCourse(db, &course )
	if result {
		return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "the data created successfully!"})
	}
	return c.JSON(fiber.Map{"error":nil,"status": "failed", "message": "the data created NOT successfully!","data":course})
}
func ActiveCourse(c *fiber.Ctx) error {
	return c.SendString("This is Post method when u want active course")
}
func CourseDetail(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	db := database.ConnectDBLocal()
	data,rows := courses.GetCourseById(db, id)
	if rows > 0 {
		return c.JSON(fiber.Map{"error":nil,"status": "failed", "message": "the data created NOT successfully!","data":data})
	}
	return c.JSON(fiber.Map{"error":"not found","status": "failed","message": "The course id "+strconv.Itoa(int(id))+" not found !"})
}
func UpdateCourse(c *fiber.Ctx) error {
	type NewC struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Price string `json:"price"`
		Description *int `json:"description", omitempty`
		} 
		
	var newCourse NewC
	test := NewC{
		ID:1,
		Title: "hello",
		Price: "300",
	}
	fmt.Println(test)
	if err := c.BodyParser(&newCourse); err != nil {
		fmt.Println("=========>",err)
		return c.JSON(fiber.Map{"error":"not found","status": false,"message": "false"})
	}
	fmt.Println(newCourse)
	
// 	return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "the data created successfully!"})

// db := database.ConnectDBLocal()
// result := courses.UpdateCourse(db, payload )
// if result {
// return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "the data created successfully!"})
// }
return c.JSON(fiber.Map{"error":nil,"status": "failed", "message": "the data created NOT successfully!","data":nil})

}
func DeleteCourse(c *fiber.Ctx) error {
	//need softdel
	id := uint32(c.QueryInt("id"))
	course := courses.Course{Id:id}
	db := database.ConnectDBLocal()
	err,rowsNum := courses.DeleteCourse(db, &course)
	if rowsNum > 0 {
		return c.JSON(fiber.Map{"error":nil,"status": "success", "message": "The course id "+strconv.Itoa(int(id))+" deleted successfully!","data":nil})	
	}
	
	return c.JSON(fiber.Map{"error":err,"status": "failed", "message": "The course id "+strconv.Itoa(int(id))+" not found !"})
}