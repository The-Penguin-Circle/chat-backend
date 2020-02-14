// Package types has all the types
package types

import "time"

type UserIdentifier string

type User struct {
	Identifier  UserIdentifier `json:"userIdentifier"`
	currentChat *Chat
}

type Chat struct {
	Users    [2]UserIdentifier `json:"users"`
	Messages []Message         `json:"messages"`
	Question string            `json:"question"`
	Answers  [2]Message        `json:"answers"`
}

type Message struct {
	SentBy  UserIdentifier `json:"sentBy"`
	At      time.Time      `json:"at"`
	Content string         `json:"content"`
}
