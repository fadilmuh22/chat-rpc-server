package main

import (
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	"github.com/fadilmuh22/chat-rpc-server/chat"
	chatv1 "github.com/fadilmuh22/chat-rpc-server/proto/chat/v1"
)

func main() {
	gs := grpc.NewServer()
	chatv1.RegisterChatServiceServer(gs, &chat.ChatService{})
	wrappedServer := grpcweb.WrapServer(gs)

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api/", wrappedServer))
	handler := cors.AllowAll().Handler(mux)

	log.Println("Serving on http://0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
