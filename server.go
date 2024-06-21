package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	chat_pb "github.com/safatanc/mesa-chat-grpc/pb/chat/proto"
	"github.com/safatanc/mesa-chat-grpc/service/chat"
	"google.golang.org/grpc"
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

	// Validator
	val := validator.New(validator.WithRequiredStructEnabled())

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
	chat_pb.RegisterChatServiceServer(server, &chat.ChatService{DB: db, Validate: val})

	log.Printf("Server is running on %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
