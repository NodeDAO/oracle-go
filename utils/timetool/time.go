package timetool

import (
	"math/rand"
	"time"
)

// RandomTime Between min Duration and max Duration, random values are taken
func RandomTime(minDuration, maxDuration time.Duration) time.Duration {
	duration := time.Duration(rand.Int63n(int64(maxDuration-minDuration+1)) + int64(minDuration))
	return duration
}
