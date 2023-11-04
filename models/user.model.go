package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string    `gorm:"type:varchar(255);not null"`
	Email string    `gorm:"uniqueIndex;not null"`
}

type CreateUserRequestPayload struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserRequestPayload struct {
	Name string `json:"name" binding:"required,min=3"`
}
