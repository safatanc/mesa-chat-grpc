package chat

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/safatanc/mesa-chat-grpc/helper"
	"github.com/safatanc/mesa-chat-grpc/model"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
	user_pb "github.com/safatanc/mesa-chat-grpc/pb/user/proto"
	"gorm.io/gorm"
)

type ChatService struct {
	chat_pb.UnimplementedChatServiceServer
	DB          *gorm.DB
	Validate    *validator.Validate
	UserService user_pb.UserServiceClient
}

// Space
func (c *ChatService) CreateSpace(ctx context.Context, request *chat_pb.CreateSpaceRequest) (*chat_pb.Space, error) {
	author, err := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
		Input: &user_pb.FindUserRequest_Id{
			Id: request.AuthorId,
		},
	})
	if err != nil {
		return nil, errors.New("author not found")
	}

	space := helper.CreateSpaceRequestToSpace(request)
	result := c.DB.Create(&space)
	if result.Error != nil {
		return nil, result.Error
	}

	return helper.SpaceToSpaceResponse(space, author), nil
}

func (c *ChatService) UpdateSpace(ctx context.Context, request *chat_pb.UpdateSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) DeleteSpace(ctx context.Context, request *chat_pb.DeleteSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) FindAllSpace(ctx context.Context, request *chat_pb.FindAllSpaceRequest) (*chat_pb.Spaces, error) {
	var spaces []*model.Space
	c.DB.Find(&spaces)

	var spaceResponses []*chat_pb.Space
	for _, space := range spaces {
		author, _ := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
			Input: &user_pb.FindUserRequest_Id{
				Id: space.AuthorID,
			},
		})
		spaceResponses = append(spaceResponses, helper.SpaceToSpaceResponse(space, author))
	}

	return &chat_pb.Spaces{
		Spaces: spaceResponses,
	}, nil
}

func (c *ChatService) FindSpace(ctx context.Context, request *chat_pb.FindSpaceRequest) (*chat_pb.Space, error) {
	var space *model.Space
	result := c.DB.First(&space, "id = ?", request.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	author, _ := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
		Input: &user_pb.FindUserRequest_Id{
			Id: space.AuthorID,
		},
	})

	return helper.SpaceToSpaceResponse(space, author), nil
}

// Message
func (c *ChatService) SendMessage(ctx context.Context, request *chat_pb.SendMessageRequest) (*chat_pb.Message, error) {
	space, err := c.FindSpace(ctx, &chat_pb.FindSpaceRequest{
		Id: request.SpaceId,
	})
	if err != nil {
		return nil, errors.New("space not found")
	}
	author, err := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
		Input: &user_pb.FindUserRequest_Id{
			Id: request.AuthorId,
		},
	})
	if err != nil {
		return nil, errors.New("author not found")
	}

	message := helper.SendMessageRequestToMessage(request)
	result := c.DB.Create(&message)

	if result.Error != nil {
		return nil, result.Error
	}

	return helper.MessageToMessageResponse(message, space, author), nil
}

func (c *ChatService) EditMessage(ctx context.Context, request *chat_pb.EditMessageRequest) (*chat_pb.Message, error) {
	message := helper.EditMessageRequestToMessage(request)
	result := c.DB.Updates(&message)
	if result.Error != nil {
		return nil, result.Error
	}

	space, _ := c.FindSpace(ctx, &chat_pb.FindSpaceRequest{
		Id: message.SpaceID,
	})
	author, _ := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
		Input: &user_pb.FindUserRequest_Id{
			Id: message.AuthorID,
		},
	})

	return helper.MessageToMessageResponse(message, space, author), nil
}

func (c *ChatService) DeleteMessage(ctx context.Context, request *chat_pb.DeleteMessageRequest) (*chat_pb.Message, error) {
	// c.DB.Delete()
	return nil, nil
}

func (c *ChatService) FindAllMessage(ctx context.Context, request *chat_pb.FindAllMessageRequest) (*chat_pb.Messages, error) {
	var messages []*model.Message
	result := c.DB.Order("created_at DESC").Find(&messages, "space_id = ?", request.SpaceId)
	if result.Error != nil {
		return nil, result.Error
	}

	var messageResponses []*chat_pb.Message
	for _, message := range messages {
		space, _ := c.FindSpace(ctx, &chat_pb.FindSpaceRequest{
			Id: message.SpaceID,
		})
		author, _ := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
			Input: &user_pb.FindUserRequest_Id{
				Id: message.AuthorID,
			},
		})
		messageResponses = append(messageResponses, helper.MessageToMessageResponse(message, space, author))
	}

	return &chat_pb.Messages{
		Messages: messageResponses,
	}, nil
}

func (c *ChatService) FindMessage(ctx context.Context, request *chat_pb.FindMessageRequest) (*chat_pb.Message, error) {
	var message *model.Message
	result := c.DB.First(&message, "id = ?", request.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	space, _ := c.FindSpace(ctx, &chat_pb.FindSpaceRequest{
		Id: request.Id,
	})
	author, _ := c.UserService.FindUser(ctx, &user_pb.FindUserRequest{
		Input: &user_pb.FindUserRequest_Id{
			Id: message.AuthorID,
		},
	})

	return helper.MessageToMessageResponse(message, space, author), nil
}
