package penguintypes

import (
	"math/rand"
)

func InsertUser(questionID int, answer string) *User {
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
	AllQueries = append(AllQueries, query)

	go FindMatches()

	return &user
}

func FindMatches() bool {
	DBMutex.Lock()
	defer DBMutex.Unlock()

	// for _, query1 := range AllQueries {
	// 	for _, query2 := range AllQueries {

	// 	}
	// }

	return true
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
