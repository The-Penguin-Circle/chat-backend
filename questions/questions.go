// Package questions keeps the question list.
package questions

import (
	"math/rand"
	"time"
)

type Question string

var questionPool []Question

func GetNQuestions(num int) []Question {
	return shuffle(questionPool)[:num]
}

func shuffle(vals []Question) []Question {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Question, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
