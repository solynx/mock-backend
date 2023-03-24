package folders

import (
	// "github.com/gofiber/fiber/v2"
	"app/models/menu/requests"
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID           uuid.UUID          `gorm:"type:char(36);primary_key;unique;not null" json:"id"`
	Name         string             `gorm:"type:varchar(100);not null" json:"name"`
	CollectionId string             `gorm:"type:char(36);not null" json:"collection_id"`
	Requests     []requests.Request `gorm:"foreignKey:FolderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"requests,omitempty"`
	CreatedAt    time.Time          `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `gorm:"not null" json:"updated_at,omitempty"`
}
