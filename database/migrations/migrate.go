package migrations

import (
	"app/models"
	"app/models/answers"
	"app/models/categories"
	"app/models/exams"
	"app/models/questions"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&exams.Exam{})
	db.AutoMigrate(&answers.Answer{})
	db.AutoMigrate(&questions.Question{})
	db.AutoMigrate(&categories.Category{})
}
