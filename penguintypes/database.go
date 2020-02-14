package penguintypes

import "sync"

// The spoof database

// DBMutex is the database mutex
var DBMutex sync.Mutex

// AllUsers contains all of the users
var AllUsers []User

// AllChats contains all of the chats
var AllChats []Chat

// AllQueries contains all of the queries
var AllQueries []ChatQuery
