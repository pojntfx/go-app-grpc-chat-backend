package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/pojntfx/go-app-grpc-chat-backend/pkg/proto/generated"
)

type ChatService struct {
	proto.ChatServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) CreateMessage(ctx context.Context, msg *proto.ChatMessage) (*proto.ChatMessage, error) {
	log.Println("Received chat message", msg)

	msg.Content = fmt.Sprintf("%v: %v", time.Now(), msg.Content)

	return msg, nil
}
