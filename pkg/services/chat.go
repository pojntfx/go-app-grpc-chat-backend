package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	proto "github.com/pojntfx/go-app-grpc-chat-backend/pkg/proto/generated"
)

type ChatService struct {
	proto.ChatServiceServer

	messageChan chan *proto.ChatMessage
}

func NewChatService() *ChatService {
	return &ChatService{messageChan: make(chan *proto.ChatMessage)}
}

func (s *ChatService) CreateMessage(ctx context.Context, msg *proto.ChatMessage) (*proto.ChatMessage, error) {
	log.Println("Received chat message", msg)

	msg.Content = fmt.Sprintf("%v: %v", time.Now(), msg.Content)

	s.messageChan <- msg

	return msg, nil
}

func (s *ChatService) SubscribeToMessages(_ *empty.Empty, stream proto.ChatService_SubscribeToMessagesServer) error {
	for message := range s.messageChan {
		if err := stream.Send(message); err != nil {
			return err
		}
	}

	return nil
}
