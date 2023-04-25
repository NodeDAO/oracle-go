package errs

import (
	"fmt"
	"time"
)

// SleepError With time. Duration's error
// The most intuitive usage is: the length of time that can be specified when encountering Sleep
type SleepError struct {
	Err   error         `json:"-"`
	Msg   string        `json:"msg"`
	Sleep time.Duration `json:"sleep"`
}

func (e *SleepError) Error() string {
	return e.Err.Error()
}

func NewSleepError(msg string, sleep time.Duration) error {
	return &SleepError{
		Err:   fmt.Errorf(msg),
		Sleep: sleep,
		Msg:   msg,
	}
}
