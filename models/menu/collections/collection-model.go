package collections

import (
	"gorm.io/gorm"
)

func GetMenu(db *gorm.DB) (collections []Collection) {
	db.Table("collections").Debug().Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Responses")
		})
	}).Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return db.Where("`requests`.`folder_id` IS NULL").Preload("Responses")
	}).Find(&collections)
	return collections
}
func CreateACollection(db *gorm.DB, collection Collection) bool {

	result := db.Create(&collection)
	return result.RowsAffected > 0
}
func GetACollection(db *gorm.DB, collection Collection) bool {

	result := db.Find(&collection)
	return result.RowsAffected > 0
}
func CheckMockApi(db *gorm.DB, collection Collection) bool {
	result := db.Find(&collection)
	if result.RowsAffected > 0 {
		return collection.IsServer
	}
	return false
}

func UpdateCollectionById(db *gorm.DB, collection Collection) bool {
	result := db.Table("collections").Debug().Where("id = ?", collection.ID).Updates(collection)

	return result.RowsAffected > 0
}
func RemoveCollection(db *gorm.DB, collection Collection) bool {
	result := db.Table("collections").Delete(collection)

	return result.RowsAffected > 0
}
