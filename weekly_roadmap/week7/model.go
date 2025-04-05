package week7

import "time"

type Message struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	SenderID   int       `json:"sender_id" gorm:"not null"`
	ReceiverID int       `json:"receiver_id" gorm:"not null"`
	Message    string    `json:"message" gorm:"type:text;not null"`
	Timestamp  time.Time `json:"timestamp" gorm:"autoCreateTime"`
}
