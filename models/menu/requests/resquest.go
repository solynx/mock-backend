package requests

import (
	// "github.com/gofiber/fiber/v2"
	"app/models/menu/responses"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID     uuid.UUID `gorm:"type:char(36);primary_key;unique;not null" json:"id,omitempty"`
	Name   string    `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	Method string    `gorm:"type:char(10);not null" json:"method,omitempty"`

	UriComponent string               `gorm:"type:varchar(255);not null" json:"uri_component"`
	Query        string               `gorm:"type:text;default:null" json:"query"`
	Header       string               `gorm:"type:text;default:null" json:"header"`
	Body         string               `gorm:"type:text;default:null" json:"body"`
	FolderId     string               `gorm:"type:char(36);default:null" json:"folder_id,omitempty"`
	CollectionId string               `gorm:"type:char(36);default:null" json:"collection_id"`
	Responses    []responses.Response `gorm:"foreignKey:RequestId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"responses"`
	CreatedAt    time.Time            `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time            `gorm:"not null" json:"updated_at,omitempty"`
}
