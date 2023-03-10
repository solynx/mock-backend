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
func GetExamById(db *gorm.DB, id int) (exam Exam) {
	db.Find(&exam, id)
	return exam
}
func UpdateExam(db *gorm.DB, exam UpdateExamStruct) bool {
	result := db.Save(exam)
	return result.RowsAffected > 0
}
func DeleteExam(db *gorm.DB, exam UpdateExamStruct) bool {
	result := db.Delete(&exam)
	return result.RowsAffected > 0
}
