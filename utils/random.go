package utils

import (
	"math/rand"
	"time"
)

func RandomNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(n)
	return a
}
