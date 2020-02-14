package types

import (
	"math/rand"
)

func insertUser(questionID int, answer string) {
	user := User{
		Identifier:  UserIdentifier(randStringRunes(10)),
		currentChat: nil,
	}
	query := ChatQuery{
		user:           &user,
		questionID:     questionID,
		questionAnswer: answer,
	}
	DBMutex.Lock()
	defer DBMutex.Unlock()
	AllUsers = append(AllUsers, user)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
