package sockets

import (
	"encoding/json"
	"errors"
	"github.com/The-Penguin-Circle/chat-backend/penguintypes"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type matchMePacket struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

func execMatchMePacket(p []byte, conn *websocket.Conn, mutex *sync.Mutex) error {
	var newPacket matchMePacket
	err := json.Unmarshal(p, &newPacket)
	if err != nil {
		return errors.New("error: could not unmarshal that json")
	}

	if newPacket.QuestionID < 0 {
		return errors.New("question id must be at least 0")
	}
	if newPacket.Answer == "" {
		return errors.New("answer cannot be empty")
	}

	user := penguintypes.InsertUser(newPacket.QuestionID, newPacket.Answer, mutex, conn)

	err = conn.WriteJSON(struct {
		Type string            `json:"type"`
		Data penguintypes.User `json:"data"`
	}{"match-me", *user})
	if err != nil {
		log.Println(err)
	}

	return nil
}
