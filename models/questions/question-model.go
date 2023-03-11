package questions

import (
	"github.com/google/uuid"
)

type Question struct {
	Id            uint32    `gorm:"primary_key;auto_increment" json:"-" query:"id"`
	ExamId        uuid.UUID `json:"-"`
	Content       string    ` json:"content, omitempty"`
	QuestionLabel string    `json:"-"`
	QuestionOrder uint32    `json:"-"`
}
