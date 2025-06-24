package services

import (
	"share-the-meal/internal/utils"
)

type WebSocketService struct {
	hub *utils.Hub
}

func NewWebSocketService(hub *utils.Hub) *WebSocketService {
	return &WebSocketService{hub: hub}
}

func (s *WebSocketService) BroadcastMessage(message []byte) {
	s.hub.Broadcast(message)
}

func (s *WebSocketService) NotifyUser(userID int64, message []byte) {
	s.hub.NotifyUser(userID, message)
}
