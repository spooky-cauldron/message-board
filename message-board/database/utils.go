package database

import (
	"log"
	"os"
	"runtime/debug"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
		debug.PrintStack()
	}
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
