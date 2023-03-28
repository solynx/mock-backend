package migrations

import (
	"app/models"
	"app/models/answers"
	"app/models/categories"
	"app/models/exams"
	"app/models/questions"

	// Collection import
	"app/models/menu/collections"
	"app/models/menu/folders"
	"app/models/menu/requests"
	"app/models/menu/responses"
	mock_server "app/models/mock-server"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&exams.Exam{})
	db.AutoMigrate(&answers.Answer{})
	db.AutoMigrate(&questions.Question{})
	db.AutoMigrate(&mock_server.MockApi{})
	db.AutoMigrate(&categories.Category{})
	db.AutoMigrate(&collections.Collection{})
	db.AutoMigrate(&folders.Folder{})
	db.AutoMigrate(&requests.Request{})
	db.AutoMigrate(&responses.Response{})
}
