// Package questions keeps the question list.
package questions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type Question struct {
	Text string `json:"text"`
	ID   int    `json:"id"`
}

var questionPool []Question

var random = rand.New(rand.NewSource(time.Now().Unix()))

func init() {
	jsonPath := "files/questions.json"

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &questionPool)
	if err != nil {
		log.Fatal(err)
	}

	for i := range questionPool {
		questionPool[i].ID = i
	}
}

func getNQuestions(num int) ([]Question, error) {
	if num > len(questionPool) {
		return nil, errors.New("too many questions requested")
	}
	return shuffle(questionPool)[:num], nil
}

func shuffle(vals []Question) []Question {
	ret := make([]Question, len(vals))
	perm := random.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
