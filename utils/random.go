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

func RandomNumExcept(n int, exceptNum int) int {
	var num int
	for {
		num := RandomNum(n)
		if num != exceptNum {
			break
		}
	}
	return num
}
