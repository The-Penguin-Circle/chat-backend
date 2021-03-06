// Package sockets handles the websocket stuff
package sockets

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type webSocketPacket struct {
	Type string `json:"type"`
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

	var mutex sync.Mutex

	go func() {
		for {
			_, p, err := conn.ReadMessage()

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

			mutex.Lock()

			switch packet.Type {
			case "match-me":
				err := execMatchMePacket(p, conn, &mutex)
				log.Println("returned answer")
				if err != nil {
					log.Println(err)
					conn.WriteMessage(1, []byte(err.Error()))
				}
			case "get-username":
				err := execGetUsername(p, conn)
				if err != nil {
					log.Println(err)
					conn.WriteMessage(1, []byte(err.Error()))
				}
			case "chat-message":
				err := execChatPacket(p, conn)
				if err != nil {
					log.Println(err)
					conn.WriteMessage(1, []byte(err.Error()))
				}
			}
			mutex.Unlock()
		}
	}()
	return nil
}
