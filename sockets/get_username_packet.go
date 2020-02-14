package sockets

import (
	"encoding/json"
	"errors"
	"github.com/The-Penguin-Circle/chat-backend/penguintypes"
	"github.com/gorilla/websocket"
)

type getUsernamePacket struct {
	GenerateNew bool   `json:"generateNew"`
	Identifier  string `json:"identifier"`
}

func execGetUsername(p []byte, conn *websocket.Conn) error {
	var newPacket getUsernamePacket
	err := json.Unmarshal(p, &newPacket)
	if err != nil {
		return errors.New("error: could not unmarshal that json")
	}

	user := penguintypes.GetUserByIdentifier(penguintypes.UserIdentifier(newPacket.Identifier))
	if user == nil {
		return errors.New("that user does not exist")
	}

	if newPacket.GenerateNew {
		user.ChangeUsername()
	}

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
