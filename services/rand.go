package services

import (
	"math/rand"
	"time"
)

// Generates random numbers
func GetRandomNumber(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(max)
}
