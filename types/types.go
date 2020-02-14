// Package types has all the types
package types

import "time"

// The UserIdentifier is a random string that is stored in the user's cookie.
type UserIdentifier string

// A User is a user.
type User struct {
	Identifier  UserIdentifier `json:"userIdentifier"`
	currentChat *Chat
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
	Users    [2]UserIdentifier `json:"users"`
	Messages []Message         `json:"messages"`
	Question string            `json:"question"`
	Answers  [2]Message        `json:"answers"`
}

// A Message is a chat maessage.
type Message struct {
	SentBy  UserIdentifier `json:"sentBy"`
	At      time.Time      `json:"at"`
	Content string         `json:"content"`
}
