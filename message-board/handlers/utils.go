package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func AssertGet(w http.ResponseWriter, r *http.Request) error {
	return AssertMethod(w, r, http.MethodGet)
}

func AssertPost(w http.ResponseWriter, r *http.Request) error {
	return AssertMethod(w, r, http.MethodPost)
}

func AssertMethod(w http.ResponseWriter, r *http.Request, method string) error {
	if r.Method != method {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}
	return nil
}

func WriteJson(data any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	dataBytes, err := json.Marshal(data)
	check(err)
	w.Write(dataBytes)
}

func Cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
