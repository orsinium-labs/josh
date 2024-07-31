package joshtest

import (
	"time"
)

// WaitFor retries try for up to maxWait.
//
// It returns nil once try returns nil the first time.
// If maxWait passes without success, it returns try's last error.
func WaitFor(maxWait time.Duration, try func() error) error {
	deadline := time.Now().Add(maxWait)
	var err error
	for time.Now().Before(deadline) {
		err = try()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	return err
}
