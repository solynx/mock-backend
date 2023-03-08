package courses

import (
	"gorm.io/gorm"
	// "database/sql/driver"
	// "encoding/json"
	// "errors"
	"fmt"
	"app/models"
)
type Courses struct {
	Id uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:255;not null" json:"description"`
	Price float32	`json:"price"`
	SalePrice float32 `json:"sale_price"`
	ImageUrl string `gorm:"size:255;not null" json:"image_url"`
	Slug string `gorm:"size:255;not null" json:"slug"`
	OwnerId uint32 `gorm:"foreignkey:OwnerId" json:"-"`
	Owner       models.User      `gorm:"foreignkey:OwnerId" json:"owner"`
	Review   Review `gorm:"foreignKey:CourseId" json:"review"`
}

type Review1 struct {
	// Id uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Star float32
	CourseId uint32 
	// UserId uint32
	Total uint32
}
// type CoursesR struct {
// 	Id uint32    `gorm:"primary_key;auto_increment" json:"id"`
// 	Title string `gorm:"size:100;not null" json="name"`
// 	Description string `gorm:"size:255;not null" json="description"`
// 	Price float32
// 	SalePrice float32
// 	ImageUrl string `gorm:"size:255;not null" json="image_url"`
// 	Slug string `gorm:"size:255;not null" json="image_url"`
// 	Review Review1
// }
// New course
// func CreateCourse (db *gorm.DB,Course *Course) (err error){
// 	err =db.Create(Course).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//Get all courses
func GetAllCourses(db *gorm.DB) (courses []Courses) {

	db.Table("courses AS C").Debug().
    Select("C.id, C.title, C.description, C.price, C.sale_price, C.image_url, C.slug, U.id AS OwnerId, U.name AS OwnerName, AVG(R.star) as Star, COUNT(R.course_id) as Total").
	
	Joins("JOIN users AS U ON U.id = C.owner_id").
	
	Joins("LEFT JOIN reviews as R ON  R.course_id=C.id").
	Preload("Review",func (db *gorm.DB) *gorm.DB{
		return db.Select("course_id, AVG(star) AS Star, COUNT(course_id) AS Total").Group("course_id")
	}).
    Preload("Owner").
    Group("C.id").
    Find(&courses)
	fmt.Println(courses)
	return courses
}
//get a course
// func GetCourseById(db *gorm.DB, Course *Course, id string) (err error) {
// 	err = db.Where("id = ?", id).First(Course).Error
// 	if err != nil {
// 	   return err
// 	}
// 	return nil
//  }
// //update a course
//  func UpdateCourse(db *gorm.DB, Course *Course) (err error) {
// 	db.Save(Course)
// 	return nil
//  }
//  // delete a course
//  func DeleteCourse(db *gorm.DB, Course *Course, id int) (err error) {
// 	db.Where("id = ?", id).Delete(Course)
// 	return nil
//  }