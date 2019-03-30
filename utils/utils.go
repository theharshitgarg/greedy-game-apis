package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomRange(min int, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
