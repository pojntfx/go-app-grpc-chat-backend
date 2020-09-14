package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"fmt"
	"log"
	"sync"
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

func (s *ChatService) TransceiveMessages(stream proto.ChatService_TransceiveMessagesServer) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				log.Println("could not receive message from client", err)

				return
			}

			s.messageChan <- message
		}
	}()

	go func() {
		for message := range s.messageChan {
			if err := stream.Send(message); err != nil {
				log.Println("could not send message to client", err)
			}
		}
	}()

	wg.Wait()

	return nil
}
