package model

import (
	"github.com/google/uuid"
	user_pb "github.com/safatanc/mesa-chat-grpc/pb/user/proto"
)

type Space struct {
	ID          uuid.UUID             `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string                `validate:"required,min=3,max=50"`
	Description string                `validate:"min=3,max=255"`
	AuthorID    string                `validate:"uuid"`
	Author      *user_pb.UserResponse `gorm:"references:id"`
	CreatedAt   int64                 `gorm:"autoCreateTime"`
	UpdatedAt   int64                 `gorm:"autoUpdateTime"`
}
