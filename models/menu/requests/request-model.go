package requests

import (
	"gorm.io/gorm"
)

func GetMenu(db *gorm.DB) (requests []Request) {
	return requests
}
func CreateARequest(db *gorm.DB, request Request) (id string) {
	result := db.Create(&request)
	if result.RowsAffected > 0 {
		return request.ID.String()
	}
	return result.Error.Error()
}
func UpdateRequestById(db *gorm.DB, request Request) bool {
	result := db.Table("requests").Debug().Where("id = ?", request.ID).Updates(request)

	return result.RowsAffected > 0
}
func RemoveRequest(db *gorm.DB, request Request) bool {
	result := db.Table("requests").Delete(request)

	return result.RowsAffected > 0
}
