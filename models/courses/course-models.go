package courses

import (
	"app/models"
	"fmt"

	"gorm.io/gorm"
)

type Course struct {
	Id          uint32      `gorm:"primary_key;auto_increment" json:"id,omitempty" query:"id"`
	Title       string      `gorm:"size:100;not null" json:"title,omitempty" query:"title"`
	Description string      `gorm:"size:255;not null" json:"description,omitempty" query:"description"`
	Price       float32     `json:"price,omitempty"`
	SalePrice   float32     `json:"sale_price,omitempty"`
	ImageUrl    string      `gorm:"size:255;not null" json:"image_url,omitempty"`
	Slug        string      `gorm:"size:255;not null" json:"slug,omitempty"`
	OwnerId     uint32      `gorm:"foreignkey:OwnerId" json:"-"`
	Owner       models.User `gorm:"foreignkey:OwnerId" json:"owner"`
	Review      Review      `gorm:"foreignKey:CourseId" json:"review"`
}

// Get all courses
func GetAllCourses(db *gorm.DB) (courses []Course, err error) {

	result := db.Table("courses AS C").
		Select("C.id, C.title, C.description, C.price, C.sale_price, C.image_url, C.slug, U.id AS OwnerId, U.name AS OwnerName, AVG(R.star) as Star, COUNT(R.course_id) as Total").
		Joins("JOIN users AS U ON U.id = C.owner_id").
		Joins("LEFT JOIN reviews as R ON  R.course_id=C.id").
		Preload("Review", func(db *gorm.DB) *gorm.DB {
			return db.Select("course_id, AVG(star) AS Star, COUNT(course_id) AS Total").Group("course_id")
		}).
		Preload("Owner").
		Group("C.id").
		Find(&courses)

	return courses, result.Error
}
func CreateCourse(db *gorm.DB, course *Course) bool {
	result := db.Create(&course)

	if result.Error != nil {
		return false
	}
	return true
}

// get a course
func GetCourseById(db *gorm.DB, id string) (course Course, rows int64) {

	result := db.Table("courses AS C").
		Select("C.id, C.title, C.description, C.price, C.sale_price, C.image_url, C.slug, U.id AS OwnerId, U.name AS OwnerName, AVG(R.star) as Star, COUNT(R.course_id) as Total").
		Debug().
		Joins("JOIN users AS U ON U.id = C.owner_id").
		Joins("LEFT JOIN reviews as R ON  R.course_id= ?", id).
		Preload("Review", func(db *gorm.DB) *gorm.DB {
			return db.Select("course_id, AVG(star) AS Star, COUNT(course_id) AS Total").Group("course_id")
		}).
		Preload("Owner").
		Where("C.id = ?", id).
		Group("C.id").
		First(&course)

	return course, result.RowsAffected
}

// //update a course
func UpdateCourse(db *gorm.DB, Course *Course) bool {
	fmt.Println(Course)
	result := db.Save(Course)

	return result.RowsAffected > 0
}

// // delete a course
func DeleteCourse(db *gorm.DB, Course *Course) (err error, rows int64) {
	//is softdel?
	result := db.Delete(&Course)

	return result.Error, result.RowsAffected
}
