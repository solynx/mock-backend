package mock_server

import (
	"time"

	"app/models/menu/collections"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MockApi struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key;unique;not null" json:"id"`

	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

// method
func NewBaseUrl(db *gorm.DB, baseUrl MockApi) bool {
	result := db.Create(&baseUrl)
	if result.RowsAffected > 0 {
		collection := collections.Collection{ID: baseUrl.ID, Name: baseUrl.Name, IsServer: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		return collections.CreateACollection(db, collection)
	}
	return false
}
func CheckBaseUrl(db *gorm.DB, baseUrl MockApi) bool {
	result := db.Find(baseUrl)
	return result.RowsAffected > 0
}
