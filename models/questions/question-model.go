package questions

import (
	// "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Question struct {
	Id            uint32    `gorm:"primary_key;auto_increment" json:"id,omitempty" query:"id"`
	ExamId        uuid.UUID `json:"exam_id,omitempty"`
	Content       string    `gorm:"not null" json:"content, omitempty"`
	QuestionLabel string    `gorm:"not null" json:"question_label, omitempty"`
	QuestionOrder uint32    `gorm:"not null" json:"question_order, omitempty"`
}
