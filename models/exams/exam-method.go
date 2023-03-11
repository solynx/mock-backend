package exams

import (
	// "github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func GetExams(db *gorm.DB) (exams []ExamResponse) {
	db.Table("exams AS E").Debug().
		Select("E.id, E.name, E.type, E.content, E.swap_answer, E.swap_question, E.status, E.user_id,E.created_at, E.updated_at, C.name AS Category, Q.content AS MetadataContent,Q.question_label AS QuestionLabel,Q.question_order AS QuestionOrder").
		Joins("JOIN categories AS C ON C.id = E.category_id").
		Joins("LEFT JOIN questions as Q ON  E.question_id= Q.id").
		Preload("Metadata", func(db *gorm.DB) *gorm.DB {
			return db.Table("questions")
		}).
		Group("E.id").
		Find(&exams)

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
