package responses

import (
	// "github.com/gofiber/fiber/v2"
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key;unique;not null" json:"id,omitempty"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	Body         string    `gorm:"type:text;default:null" json:"body, omitempty"`
	Method       string    `gorm:"type:char(10);not null" json:"method,omitempty"`
	UriComponent string    `gorm:"type:varchar(255);not null" json:"uri_component,omitempty"`
	RequestId    string    `gorm:"type:char(36);not null" json:"request_id,omitempty"`

	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
