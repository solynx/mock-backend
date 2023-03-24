package responses

import (
	"gorm.io/gorm"
)

func GetMenu(db *gorm.DB) (responses []Response) {
	return responses
}
func CreateAResponse(db *gorm.DB, response Response) bool {
	result := db.Create(&response)
	return result.RowsAffected > 0

}
func UpdateResponseById(db *gorm.DB, response Response) bool {
	result := db.Table("responses").Debug().Where("id = ?", response.ID).Updates(response)

	return result.RowsAffected > 0
}
func RemoveResponse(db *gorm.DB, response Response) bool {
	result := db.Table("responses").Delete(response)

	return result.RowsAffected > 0
}
