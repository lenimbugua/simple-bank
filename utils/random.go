package utils

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt returns a random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString returns a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

//RandomAccount returns a random owner name
func RandomName() string {
	return RandomString(6)
}

//RandomCurrency returns a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	n := rand.Intn(len(currencies))
	return currencies[n]
}

func RandomAmount() int64 {
	return RandomInt(1, 1000)
}
