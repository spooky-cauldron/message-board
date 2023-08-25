package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJson(data any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	dataBytes, err := json.Marshal(data)
	check(err)
	w.Write(dataBytes)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
