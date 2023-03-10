package exams

import (
	// "github.com/gofiber/fiber/v2"
	"time"

	"github.com/google/uuid"
)

type Exam struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name, omitempty"`
	Type         string    `gorm:"type:varchar(50);not null" json:"type, omitempty"`
	CategoryId   uint32    `json:"category_id, omitempty"`
	Content      string    `gorm:"type:varchar(255);not null" json:"content, omitempty"`
	SwapAnswer   uint32    `json:"swap_anwser,omitempty"`
	SwapQuestion uint32    `json:"swap_question,omitempty"`
	Status       string    `gorm:"type:varchar(255);not null" json:"status,omitempty"`
	UserId       uuid.UUID `gorm:"type:char(36) ;not null" json:"user_id,omitempty"`
	QuestionId   uint32    `json:"question_id,omitempty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateExamStruct struct {
	Id           uuid.UUID
	Name         string `json:"name"  binding:"required"`
	Type         string `json:"type"  binding:"required"`
	CategoryId   uint32 `json:"category_id"  binding:"required"`
	Content      string `json:"content"  binding:"required"`
	SwapAnswer   uint32 `json:"swap_anwser,omitempty"`
	SwapQuestion uint32 `json:"swap_question,omitempty"`
	Status       string `json:"status"  binding:"required"`
	UserId       uuid.UUID
	QuestionId   uint32 `json:"question_id,omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type UpdateExamStruct struct {
	Name         string    `json:"name,omitempty"`
	Type         string    `json:"type,omitempty"`
	CategoryId   uint32    `json:"category_id,omitempty"`
	Content      string    `json:"content,omitempty"`
	SwapAnswer   uint32    `json:"swap_anwser,omitempty"`
	SwapQuestion uint32    `json:"swap_question,omitempty"`
	Status       string    `json:"status,omitempty"`
	UserId       uuid.UUID `json:"user_id,omitempty"`
	QuestionId   uint32    `json:"question_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

//exams response
