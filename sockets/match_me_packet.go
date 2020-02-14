package sockets

import (
	"encoding/json"
	"errors"
	"github.com/The-Penguin-Circle/chat-backend/penguintypes"
	"github.com/gorilla/websocket"
)

type matchMePacket struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

func execMatchMePacket(p []byte, conn *websocket.Conn) error {
	var newPacket matchMePacket
	err := json.Unmarshal(p, &newPacket)
	if err != nil {
		return errors.New("error: could not unmarshal that json")
	}

	if newPacket.QuestionID < -1 {
		return errors.New("question id must be at least -1")
	}
	if newPacket.Answer == "" {
		return errors.New("answer cannot be empty")
	}

	user := penguintypes.InsertUser(newPacket.QuestionID, newPacket.Answer)

	user.WebSocket = conn

	responseInBytes, err := json.Marshal(user)

	if err != nil {
		return err
	}

	err = conn.WriteMessage(1, responseInBytes)
	if err != nil {
		return err
	}

	return nil
}
