package requests

import (
	// "github.com/gofiber/fiber/v2"
	"app/models/menu/responses"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Id     uuid.UUID `gorm:"type:char(36);primary_key;unique;not null" json:"id,omitempty"`
	Name   string    `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	Method string    `gorm:"type:char(10);not null" json:"method,omitempty"`

	UriComponent string               `gorm:"type:varchar(255);not null" json:"uri_component,omitempty"`
	FolderId     string               `gorm:"type:char(36);" json:"folder_id,omitempty"`
	CollectionId string               `gorm:"type:char(36);not null" json:"collection_id,omitempty"`
	Responses    []responses.Response `gorm:"references:RequestId" json:"responses,omitempty"`
	CreatedAt    time.Time            `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time            `gorm:"not null" json:"updated_at,omitempty"`
}
