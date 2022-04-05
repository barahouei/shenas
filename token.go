package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

//This function generate an unique and reusable token which can be used as user's custom link
//or a security approch if it was needed.
func token() string {
	crutime := time.Now().Unix()
	hash := md5.New()

	io.WriteString(hash, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", hash.Sum(nil))

	return token
}
