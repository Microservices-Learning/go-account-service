package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email    string    `json:"email" gorm:"varchar(255);not null"`
	Password string    `json:"password" gorm:"varchar(255);not null"`
}
