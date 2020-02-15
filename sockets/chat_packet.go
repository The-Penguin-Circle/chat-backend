package sockets

import (
	"encoding/json"
	"errors"
	"github.com/The-Penguin-Circle/chat-backend/penguintypes"
	"github.com/gorilla/websocket"
	"log"
)

type chatPacket struct {
	Identifier penguintypes.UserIdentifier `json:"identifier"`
	Message    string                      `json:"message"`
}

func execChatPacket(p []byte, conn *websocket.Conn) error {
	var newPacket chatPacket
	err := json.Unmarshal(p, &newPacket)
	if err != nil {
		return errors.New("error: could not unmarshal that json")
	}

	if newPacket.Message == "" {
		return errors.New("message cannot be empty")
	}

	user := penguintypes.GetUserByIdentifier(newPacket.Identifier)
	if user == nil {
		return errors.New("that user does not exist")
	}

	var otherUser penguintypes.User
	if user.CurrentChat == nil {
		return errors.New("no such chat")
	}
	for _, u := range user.CurrentChat.Users {
		if user.Identifier != u.Identifier {
			otherUser = u
			break
		}
	}

	type messagePacket struct {
		Message string `json:"message"`
	}

	go func() {
		otherUser.WebSocketMutex.Lock()
		defer otherUser.WebSocketMutex.Unlock()

		err = otherUser.WebSocket.WriteJSON(struct {
			Type string        `json:"type"`
			Data messagePacket `json:"data"`
		}{
			Type: "chat-message",
			Data: messagePacket{
				Message: newPacket.Message,
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	err = conn.WriteJSON(struct {
		Type string `json:"type"`
	}{"ok"})
	if err != nil {
		log.Println(err)
	}

	return nil
}
