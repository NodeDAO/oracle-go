package timetool

import (
	"math/rand"
	"time"
)

// SleepWithRandom sleeps for a duration between minDuration and maxDuration, inclusive
// duration is randomly selected between minDuration and maxDuration
func SleepWithRandom(minDuration, maxDuration time.Duration) {
	duration := time.Duration(rand.Int63n(int64(maxDuration-minDuration+1)) + int64(minDuration))
	time.Sleep(duration)
}
