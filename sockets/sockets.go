// Package sockets handles the websocket stuff
package sockets

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "websocket could not be opened", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	_ = conn
}
