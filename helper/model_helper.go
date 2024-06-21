package helper

import (
	"github.com/safatanc/mesa-chat-grpc/model"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
)

func SpaceToSpaceResponse(space *model.Space) *chat_pb.Space {
	return &chat_pb.Space{
		Id:          space.ID.String(),
		Title:       space.Title,
		Description: space.Description,
		CreatedAt:   space.CreatedAt,
		UpdatedAt:   space.UpdatedAt,
	}
}
