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
	space.Author = author
	result := c.DB.Create(&space)
	if result.Error != nil {
		return nil, result.Error
	}

	return helper.SpaceToSpaceResponse(space), nil
}

func (c *ChatService) UpdateSpace(ctx context.Context, request *chat_pb.UpdateSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) DeleteSpace(ctx context.Context, request *chat_pb.DeleteSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) FindAllSpace(ctx context.Context, request *chat_pb.FindAllSpaceRequest) (*chat_pb.Spaces, error) {
	var spaces []*model.Space
	c.DB.Joins("Author").Find(&spaces)

	var spaceResponses []*chat_pb.Space
	for _, space := range spaces {
		spaceResponses = append(spaceResponses, helper.SpaceToSpaceResponse(space))
	}

	return &chat_pb.Spaces{
		Spaces: spaceResponses,
	}, nil
}

func (c *ChatService) FindSpace(ctx context.Context, request *chat_pb.FindSpaceRequest) (*chat_pb.Space, error) {
	var space *model.Space
	result := c.DB.Joins("Author").First(&space, "spaces.id = ?", request.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return helper.SpaceToSpaceResponse(space), nil
}

// Message
func (c *ChatService) SendMessage(ctx context.Context, request *chat_pb.SendMessageRequest) (*chat_pb.Message, error) {
	return nil, nil
}

func (c *ChatService) EditMessage(ctx context.Context, request *chat_pb.EditMessageRequest) (*chat_pb.Message, error) {
	return nil, nil
}

func (c *ChatService) DeleteMessage(ctx context.Context, request *chat_pb.DeleteMessageRequest) (*chat_pb.Message, error) {
	return nil, nil
}
