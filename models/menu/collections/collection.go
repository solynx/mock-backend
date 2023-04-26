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
	Name      string             `gorm:"type:varchar(100);not null" json:"name"`
	UserId    uint32             `gorm:"not null;default:1" json:"user_id"`
	Requests  []requests.Request `gorm:"foreignKey:CollectionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"requests"`
	Folders   []folders.Folder   `gorm:"foreignKey:CollectionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"folders"`
	IsServer  bool               `gorm:"not null;default:false" json:"is_server"`
	Variable  string             `gorm:"type:text;default:null" json:"variable"`
	CreatedAt time.Time          `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time          `gorm:"not null" json:"updated_at,omitempty"`
}
