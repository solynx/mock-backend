package responses

import (
	"gorm.io/gorm"
)

func GetMenu(db *gorm.DB) (responses []Response) {
	return responses
}
func CreateAResponse(db *gorm.DB, response Response) (id string) {
	result := db.Create(&response)
	if result.RowsAffected > 0 {
		return response.ID.String()
	}
	return result.Error.Error()
}
func UpdateResponseById(db *gorm.DB, response Response) bool {
	result := db.Table("responses").Debug().Where("id = ?", response.ID).Updates(response)

	return result.RowsAffected > 0
}
func RemoveResponse(db *gorm.DB, response Response) bool {
	result := db.Table("responses").Delete(response)

	return result.RowsAffected > 0
}
