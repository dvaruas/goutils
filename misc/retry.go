package miscutils

import (
	"fmt"
	"time"
)

func RetryFunc(
	timeout time.Duration,
	fn func() error,
) error {
	var err error
	deadline := time.Now().Add(timeout)
	// Initial time to sleep between tries.
	pause := 50 * time.Millisecond
	// Cutoff for exponential backoff.
	maxPause := 1 * time.Second
	for tryCount := 0; time.Until(deadline) >= 0; {
		if err = fn(); err == nil {
			return nil
		}

		time.Sleep(pause)
		pause = 2 * pause
		if pause > maxPause {
			pause = maxPause
		}
		tryCount++
		fmt.Printf("RetryFunc: try [%v], error: %v\n", tryCount, err)
	}
	return err
}
