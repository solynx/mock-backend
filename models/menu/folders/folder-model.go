package folders

import (
	"gorm.io/gorm"
)

func GetMenu(db *gorm.DB) (folders []Folder) {
	return folders
}
func CreateAFolder(db *gorm.DB, folder Folder) (id string) {
	result := db.Create(&folder)
	if result.RowsAffected > 0 {
		return folder.ID.String()
	}
	return result.Error.Error()
}
func UpdateFolderById(db *gorm.DB, folder Folder) bool {
	result := db.Table("folders").Debug().Where("id = ?", folder.ID).Updates(folder)

	return result.RowsAffected > 0
}
func RemoveFolder(db *gorm.DB, folder Folder) bool {
	result := db.Table("folders").Delete(folder)

	return result.RowsAffected > 0
}
