package helper

import (
	"github.com/safatanc/mesa-chat-grpc/model"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
	user_pb "github.com/safatanc/mesa-chat-grpc/pb/user/proto"
)

func CreateSpaceRequestToSpace(request *chat_pb.CreateSpaceRequest) *model.Space {
	return &model.Space{
		Title:       request.Title,
		Description: request.Description,
		AuthorID:    request.AuthorId,
	}
}

func SpaceToSpaceResponse(space *model.Space, author *user_pb.UserResponse) *chat_pb.Space {
	return &chat_pb.Space{
		Id:          space.ID.String(),
		Title:       space.Title,
		Description: space.Description,
		Author:      author,
		CreatedAt:   space.CreatedAt,
		UpdatedAt:   space.UpdatedAt,
	}
}

func SendMessageRequestToMessage(request *chat_pb.SendMessageRequest) *model.Message {
	return &model.Message{
		SpaceID:  request.SpaceId,
		AuthorID: request.AuthorId,
		Content:  request.Content,
	}
}

func EditMessageRequestToMessage(request *chat_pb.EditMessageRequest) *model.Message {
	return &model.Message{
		Content: request.Content,
	}
}

func MessageToMessageResponse(message *model.Message, space *chat_pb.Space, author *user_pb.UserResponse) *chat_pb.Message {
	return &chat_pb.Message{
		Id:        message.ID.String(),
		SpaceId:   message.SpaceID,
		Space:     space,
		AuthorId:  message.AuthorID,
		Author:    author,
		Content:   message.Content,
		CreatedAt: space.CreatedAt,
		UpdatedAt: space.UpdatedAt,
	}
}
