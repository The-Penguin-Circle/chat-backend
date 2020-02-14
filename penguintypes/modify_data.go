package penguintypes

import (
	"github.com/The-Penguin-Circle/chat-backend/generation"
	"math/rand"
)

func InsertUser(questionID int, answer string) *User {
	user := User{
		Identifier:   UserIdentifier(randStringRunes(10)),
		Username:     generation.GenerateUsername(),
		ProfileImage: generation.GenerateImage(),
	}
	query := ChatQuery{
		user:           &user,
		questionID:     questionID,
		questionAnswer: answer,
	}
	dbMutex.Lock()
	AllUsers = append(AllUsers, user)
	AllQueries = append(AllQueries, query)
	dbMutex.Unlock()

	go findMatches()

	return &user
}

func (u *User) ChangeUsername() {
	u.ProfileImage = generation.GenerateImage()
	u.Username = generation.GenerateUsername()
}

func findMatches() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	remove := func(q ChatQuery) {
		for i, q3 := range AllQueries {
			if q3 == q {
				AllQueries = append(AllQueries[:i], AllQueries[i+1:]...)
			}
		}
	}

	onSuccess := func(query1 ChatQuery, query2 ChatQuery) Chat {
		remove(query1)
		remove(query2)
		newChat := Chat{
			Users: [2]User{
				*query1.user,
				*query2.user,
			},
			QuestionID: query1.questionID,
			Answers: [2]string{
				query1.questionAnswer,
				query2.questionAnswer,
			},
		}
		AllChats = append(AllChats, newChat)

		type matchResponse struct {
			OtherUser     User   `json:"otherUser"`
			OtherResponse string `json:"otherResponse"`
		}

		query1.user.WebSocket.WriteJSON(
			struct {
				Type string        `json:"type"`
				Data matchResponse `json:"data"`
			}{"chat-found", matchResponse{
				OtherUser:     *query2.user,
				OtherResponse: query2.questionAnswer,
			}},
		)
		query2.user.WebSocket.WriteJSON(
			struct {
				Type string        `json:"type"`
				Data matchResponse `json:"data"`
			}{"chat-found", matchResponse{
				OtherUser:     *query1.user,
				OtherResponse: query1.questionAnswer,
			}},
		)

		return newChat
	}

	for _, query1 := range AllQueries {
		for _, query2 := range AllQueries {
			if query1 != query2 && query1.questionID == query2.questionID {
				identifer1 := query1.user.Identifier
				identifer2 := query2.user.Identifier
				chat := onSuccess(query1, query2)
				setChat := func(identifier UserIdentifier) {
					for i := range AllUsers {
						if AllUsers[i].Identifier == identifier {
							AllUsers[i].CurrentChat = &chat
						}
					}
				}

				setChat(identifer1)
				setChat(identifer2)
				return
			}
		}
	}
}

func GetUserByIdentifier(identifier UserIdentifier) *User {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	for _, u := range AllUsers {
		if u.Identifier == identifier {
			return &u
		}
	}
	return nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
