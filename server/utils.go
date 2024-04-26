package main

import (
	"math/rand"
	"time"
)

const nameCodeSize = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	randomString := make([]byte, nameCodeSize)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
