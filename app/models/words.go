package models

type Words struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	From    string `gorm:"size:255" json:"from"`
	Message string `gorm:"size:5000" json:"message"`
}
