package server

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStr(runes []rune, length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = runes[rand.Intn(len(runes))]
	}
	return string(s)
}
