package courses

import (
	// "gorm.io/gorm"

)
// type Course struct {
// 	Id uint32    `gorm:"primary_key;auto_increment" json:"id"`
// 	Title string `gorm:"size:100;not null" json="name"`
// 	Description string `gorm:"size:255;not null" json="description"`
// 	Price float32
// 	SalePrice float32
// 	ImageUrl string `gorm:"size:255;not null" json="image_url"`
// 	Slug string `gorm:"size:255;not null" json="image_url"`
// 	OwnerId uint32
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
// func GetCourses(db *gorm.DB,Course []*Course) (err error) {
// 	err = db.Find(Course).Error
// 	if err != nil {
// 		return err
// 	 }
// 	 return nil
// }
// //get a course
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