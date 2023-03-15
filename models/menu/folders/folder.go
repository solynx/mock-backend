package folders

import (
	// "github.com/gofiber/fiber/v2"
	"app/models/menu/requests"
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	Id           uuid.UUID          `gorm:"type:char(36);primary_key;unique;not null" json:"id,omitempty"`
	Name         string             `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	CollectionId string             `gorm:"type:char(36);not null" json:"collection_id,omitempty"`
	Requests     []requests.Request `gorm:"references:FolderId" json:"requests,omitempty"`
	CreatedAt    time.Time          `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `gorm:"not null" json:"updated_at,omitempty"`
}
