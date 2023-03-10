package exams

import (
	// "github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func GetExams(db *gorm.DB) (exams []Exam) {
	db.Find(&exams)
	return exams
}
func CreateExam(db *gorm.DB, exam CreateExamStruct) (err error) {
	result := db.Table("exams").Create(&exam)
	return result.Error
}
func GetExamById(db *gorm.DB, id string) (exam Exam, rows int64) {
	result := db.Where("id = ?", id).Find(&exam)
	return exam, result.RowsAffected
}
func UpdateExam(db *gorm.DB, exam UpdateExamStruct) bool {
	result := db.Table("exams").Where("id = ?", exam.Id).Updates(exam)
	return result.RowsAffected > 0
}
func DeleteExam(db *gorm.DB, exam DeleteExamStruct) bool {
	result := db.Table("exams").Delete(&exam)
	return result.RowsAffected > 0
}
