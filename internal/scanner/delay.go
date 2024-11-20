package scanner

import (
	"math/rand"
	"time"
)

func AddRandomDelay() {
	minDelay := 1  // Minimum delay in seconds
	maxDelay := 5  // Maximum delay in seconds
	time.Sleep(time.Duration(rand.Intn(maxDelay-minDelay)+minDelay) * time.Second)
}
