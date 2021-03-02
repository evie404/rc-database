package testhelpers

import (
	"math/rand"
	"time"
)

const alphanumericBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandAlphanumericString(n int) string {
	return string(RandAlphanumericBytes(n))
}

func RandAlphanumericBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumericBytes[rand.Intn(len(alphanumericBytes))]
	}
	return b
}
