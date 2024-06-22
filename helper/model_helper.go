package helper

import (
	"github.com/safatanc/mesa-chat-grpc/model"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
)

func CreateSpaceRequestToSpace(request *chat_pb.CreateSpaceRequest) *model.Space {
	return &model.Space{
		Title:       request.Title,
		Description: request.Description,
		AuthorID:    request.AuthorId,
	}
}

func SpaceToSpaceResponse(space *model.Space) *chat_pb.Space {
	return &chat_pb.Space{
		Id:          space.ID.String(),
		Title:       space.Title,
		Description: space.Description,
		Author:      space.Author,
		CreatedAt:   space.CreatedAt,
		UpdatedAt:   space.UpdatedAt,
	}
}
