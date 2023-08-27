package handlers

import (
	"log"
	"net/http"
)

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
