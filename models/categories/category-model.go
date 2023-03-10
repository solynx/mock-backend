package categories

type Category struct {
	Id          uint32 `gorm:"primary_key;auto_increment" json:"id,omitempty" query:"id"`
	Name        string `gorm:"not null" json:"name,omitempty"`
	Description string `json:"description, omitempty"`
}
