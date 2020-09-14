package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"log"
	"sync"

	proto "github.com/pojntfx/go-app-grpc-chat-backend/pkg/proto/generated"
	"github.com/ugjka/messenger"
)

type ChatService struct {
	proto.ChatServiceServer

	messenger *messenger.Messenger
}

func NewChatService(messenger *messenger.Messenger) *ChatService {
	return &ChatService{messenger: messenger}
}

func (s *ChatService) TransceiveMessages(stream proto.ChatService_TransceiveMessagesServer) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				log.Println("could not receive message")

				break
			}

			s.messenger.Broadcast(message)
		}

		wg.Done()
	}()

	go func() {
		client, err := s.messenger.Sub()
		if err != nil {
			log.Println("could not subscribe to messages")

			s.messenger.Unsub(client)

			wg.Done()

			return
		}

		for message := range client {
			if err := stream.Send(message.(*proto.ChatMessage)); err != nil {
				log.Println("could not send message")

				break
			}
		}

		s.messenger.Unsub(client)

		wg.Done()
	}()

	wg.Wait()

	return nil
}
