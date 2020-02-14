package penguintypes

import "sync"

// The spoof database

// dbMutex is the database mutex
var dbMutex sync.Mutex

// AllUsers contains all of the users
var AllUsers []User

// AllChats contains all of the chats
var AllChats []Chat

// AllQueries contains all of the queries
var AllQueries []ChatQuery
