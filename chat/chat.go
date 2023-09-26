package chat

import (
	"context"
	"fmt"
	"sync"

	chatv1 "github.com/fadilmuh22/chat-rpc-server/proto/chat/v1"
)

type ChatService struct {
	chatv1.UnimplementedChatServiceServer

	userList  []*chatv1.User
	userMutex sync.Mutex
	clients   map[string]chan *chatv1.ChatMessage
}

func (s *ChatService) Join(ctx context.Context, user *chatv1.User) (*chatv1.JoinResponse, error) {
	fmt.Printf("User %s joined\n", user.Id)
	s.userMutex.Lock()
	defer s.userMutex.Unlock()

	// Initialize the clients map if it's nil.
	if s.clients == nil || len(s.clients) == 0 {
		s.clients = make(map[string]chan *chatv1.ChatMessage)
	}

	// Add the user to the user list.
	s.userList = append(s.userList, user)

	// Create a channel for the user's messages.
	s.clients[user.Id] = make(chan *chatv1.ChatMessage)

	return &chatv1.JoinResponse{
		Error: 0,
		Msg:   "Joined successfully",
	}, nil
}

func (s *ChatService) SendMsg(ctx context.Context, msg *chatv1.ChatMessage) (*chatv1.Empty, error) {
	fmt.Printf("Received message from %s: %s\n", msg.From, msg.Msg)
	// Broadcast the message to all connected clients.
	for _, ch := range s.clients {
		ch <- msg
	}

	fmt.Println("Broadcasted message to all clients")

	return &chatv1.Empty{}, nil
}

func (s *ChatService) ReceiveMsg(req *chatv1.ReceiveMsgRequest, stream chatv1.ChatService_ReceiveMsgServer) error {
	fmt.Printf("User %s connected\n", req.User)
	// Create a channel for the client's messages.
	clientChan := s.clients[req.User]

	for {
		select {
		case msg := <-clientChan:
			// Send the received message to the client.
			if err := stream.Send(msg); err != nil {
				return err
			}
		case <-stream.Context().Done():
			// Client disconnected, remove the channel and exit.
			delete(s.clients, req.User)
			return stream.Context().Err()
		}
	}
}

func (s *ChatService) GetAllUsers(ctx context.Context, empty *chatv1.Empty) (*chatv1.UserList, error) {
	fmt.Println("Received GetAllUsers request", s.userList)
	s.userMutex.Lock()
	defer s.userMutex.Unlock()

	return &chatv1.UserList{
		Users: s.userList,
	}, nil
}
