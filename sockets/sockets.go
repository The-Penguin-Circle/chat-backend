// Package sockets handles the websocket stuff
package sockets

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type webSocketPacket struct {
	Type string `json:"type"`
}

type matchMePacket struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin is the very safe origin function
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return errors.New("websocket could not be opened")
	}

	go func() {
		for {
			messageType, p, err := conn.ReadMessage()

			log.Println(messageType)
			log.Println(p)
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}

			var packet webSocketPacket

			err = json.Unmarshal(p, &packet)
			if err != nil {
				conn.WriteMessage(1, []byte("error: could not unmarshal that json"))
				log.Println(err)
			}

			switch packet.Type {
			case "match-me":
				err := execMatchMePacket(p, conn)
				if err != nil {
					log.Println(err)
					conn.WriteMessage(1, []byte(err.Error()))
				}
			}

		}
	}()
	return nil
}
