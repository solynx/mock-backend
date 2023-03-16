package collections

import (
	// "github.com/gofiber/fiber/v2"
	"app/models/menu/folders"
	"app/models/menu/requests"
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID        uuid.UUID          `gorm:"type:char(36);primary_key;unique;not null" json:"id"`
	Name      string             `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	UserId    uint32             `gorm:"not null" json:"user_id,omitempty"`
	Requests  []requests.Request `gorm:"foreignKey:CollectionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"requests,omitempty"`
	Folders   []folders.Folder   `gorm:"foreignKey:CollectionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"folders,omitempty"`
	CreatedAt time.Time          `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time          `gorm:"not null" json:"updated_at,omitempty"`
}
