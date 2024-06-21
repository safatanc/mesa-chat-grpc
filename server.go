package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/safatanc/mesa-chat-grpc/model"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
	user_pb "github.com/safatanc/mesa-chat-grpc/pb/user/proto"
	"github.com/safatanc/mesa-chat-grpc/service/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()

	// Database
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Space{}, &model.Message{})

	// Validator
	val := validator.New(validator.WithRequiredStructEnabled())

	// gRPC Client
	conn, err := grpc.NewClient(os.Getenv("USER_GRPC_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userClient := user_pb.NewUserServiceClient(conn)

	// gRPC Server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	chat_pb.RegisterChatServiceServer(server, &chat.ChatService{
		DB:          db,
		Validate:    val,
		UserService: userClient,
	})

	log.Printf("Server is running on %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
