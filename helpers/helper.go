package helpers

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
)

const (
	colorRed string = "\033[31m"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func LoggerError(err error) {
	if err != nil {
		fmt.Println(colorRed)
		log.Println(colorRed, "========== Start Error Message ==========")
		log.Println(colorRed, "ERROR => "+err.Error()+".")
		log.Println(colorRed, "========== End Of Error Message ==========")
	}
}

func CreateKey(email, emailkey string) (key string) {

	data := email + "|" + emailkey
	key = base64.StdEncoding.EncodeToString([]byte(data))

	return
}
