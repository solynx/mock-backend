package answers

type Answer struct {
	Id         uint32 `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	QuestionId uint32 `json:"question_id,omitempty"`
	Content    string `gorm:"not null" json:"content, omitempty"`
	IsCorrect  bool   `json:"is_correct,omitempty"`
}
