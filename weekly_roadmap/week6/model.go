package week6

type User struct {
    ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Username string `json:"username" gorm:"type:varchar(100);unique;not null"`
    Password string `json:"password" gorm:"type:varchar(100);not null"`
}