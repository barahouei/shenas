package main

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//This function creates a random string which contains character and number and returns that random string.
func randomString(length int) string {
	s := make([]byte, length)

	for i := range s {
		s[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(s)
}

//This function generate an unique and reusable token which can be used as user's custom link
func token() string {
	token := randomString(8)

	return token
}
