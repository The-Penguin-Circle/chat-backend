// Package penguintypes has all the types
package penguintypes

import "github.com/gorilla/websocket"

import "sync"

// The UserIdentifier is a random string that is stored in the user's cookie.
type UserIdentifier string

// A User is a user.
type User struct {
	Identifier     UserIdentifier  `json:"identifier"`
	Username       string          `json:"username"`
	ProfileImage   string          `json:"image"`
	WebSocket      *websocket.Conn `json:"-"`
	WebSocketMutex *sync.Mutex     `json:"-"`
	CurrentChat    *Chat           `json:"-"`
}

// A ChatQuery is the query that is stored in the database as long as
// the user has not found a partner yet.
type ChatQuery struct {
	user           *User
	questionID     int
	questionAnswer string
}

// A Chat is a chat between two users with messages and answers of questions.
type Chat struct {
	Users      [2]User   `json:"users"`
	QuestionID int       `json:"questionID"`
	Answers    [2]string `json:"answers"`
}
