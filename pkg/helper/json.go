package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(decoder); err != nil {
		return errors.New("failed decode json")
	}
	return nil
}

func WriteToResponBody(w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		return errors.New("failed to write body")
	}
	return nil
}