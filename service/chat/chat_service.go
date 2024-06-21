package chat

import (
	"context"

	"github.com/go-playground/validator/v10"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
	"gorm.io/gorm"
)

type ChatService struct {
	chat_pb.UnimplementedChatServiceServer
	DB       *gorm.DB
	Validate *validator.Validate
}

// Space
func (c *ChatService) CreateSpace(ctx context.Context, request *chat_pb.CreateSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) UpdateSpace(ctx context.Context, request *chat_pb.UpdateSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) DeleteSpace(ctx context.Context, request *chat_pb.DeleteSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
}

func (c *ChatService) FindAllSpace(ctx context.Context, request *chat_pb.FindAllSpaceRequest) (*chat_pb.Spaces, error) {
	return nil, nil
}

func (c *ChatService) FindSpace(ctx context.Context, request *chat_pb.FindSpaceRequest) (*chat_pb.Space, error) {
	return nil, nil
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
