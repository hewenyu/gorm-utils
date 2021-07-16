package utils

import "math/rand"

func RandomString(n int) string {

	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	length := len(letter)

	for i := range b {
		b[i] = letter[rand.Intn(length)]
	}

	return string(b)
}

func NonceStr() string {

	return RandomString(32)
}
