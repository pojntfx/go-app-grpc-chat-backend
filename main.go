package main

import (
	"flag"
	"log"
	"net"
	"sync"

	proto "github.com/pojntfx/go-app-grpc-chat-backend/pkg/proto/generated"
	"github.com/pojntfx/go-app-grpc-chat-backend/pkg/services"
	"github.com/pojntfx/go-app-grpc-chat-backend/pkg/websocketproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	laddr := flag.String("laddr", "0.0.0.0:1925", "Listen address.")
	wsladdr := flag.String("wsladdr", "0.0.0.0:10000", "Listen address (for the WebSocket proxy).")

	flag.Parse()

	lis, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	proxy := websocketproxy.NewWebSocketProxyServer(*wsladdr)
	wslis, err := proxy.Listen()
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	chatService := services.NewChatService()
	proto.RegisterChatServiceServer(srv, chatService)

	log.Printf("Listening on %v and %v (WebSocket proxy)", *laddr, *wsladdr)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	go func() {
		if err := srv.Serve(wslis); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()
}
