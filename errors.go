package main

import "log"

// This is a function which deals with the errors.
func errorChecking(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
