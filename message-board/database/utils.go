package database

import (
	"log"
	"runtime/debug"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
		debug.PrintStack()
	}
}
