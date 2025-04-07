package models

type Student struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Age   int    `json:"age" gorm:"not null"`
	Class string `json:"class" gorm:"type:varchar(50)"`
}
