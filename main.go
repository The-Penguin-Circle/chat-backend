package main

import (
	"github.com/The-Penguin-Circle/chat-backend/questions"
	"github.com/The-Penguin-Circle/chat-backend/sockets"
	"net/http"

	"log"
)

func errorClosure(toCall func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := toCall(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	http.HandleFunc("/get-questions", errorClosure(questions.ServeQuestions))
	http.HandleFunc("/websocket", errorClosure(sockets.WebSocket))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
