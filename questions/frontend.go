package questions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func ServeQuestions(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("must use GET method")
	}
	query := r.URL.Query()
	numberString := query.Get("n")
	number, err := strconv.Atoi(numberString)
	if err != nil {
		return errors.New("must set ?n=, as a number")
	}
	result, err := getNQuestions(number)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return errors.New("error while creating JSON, this should be our fault")
	}
	w.Write(bytes)
	return nil
}
