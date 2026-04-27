package random

import "math/rand/v2"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomByte() byte {
	return charset[rand.IntN(len(charset))]
}

func NewRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomByte()
	}
	return string(b)
}
