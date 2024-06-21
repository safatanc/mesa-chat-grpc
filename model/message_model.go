package model

import "github.com/google/uuid"

type Message struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SpaceID   string    `validate:"required,uuid"`
	AuthorID  string    `validate:"required,uuid"`
	Space     *Space
	Content   string `validate:"required"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}
