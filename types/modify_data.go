package types

import (
	"math/rand"
)

func insertUser(questionID int, answer string) *User {
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
	AllUsers = append(AllUsers, user)
	AllQueries = append(AllQueries, query)
	DBMutex.Unlock()
	go findMatches()

	return &user
}

func findMatches() {
	DBMutex.Lock()
	defer DBMutex.Unlock()
	// TODO: Find matches and add resulting chat to database, remove the query from AllQueries
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
