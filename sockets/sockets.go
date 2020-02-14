// Package sockets handles the websocket stuff
package sockets

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type webSocketPacket struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
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
			}

			var packet 

			err = conn.WriteMessage(messageType, p)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	return nil
}
