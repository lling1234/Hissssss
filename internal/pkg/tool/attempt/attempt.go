package attempt

import (
	"math/rand"
	"time"
)

type StopErr struct {
	error
}

func StopRetry(err error) StopErr {
	return StopErr{err}
}

func Attempt(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if s, ok := err.(StopErr); ok {
			return s.error
		}
		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return Attempt(attempts, 2*sleep, f)
		}
		return err
	}
	return nil
}
