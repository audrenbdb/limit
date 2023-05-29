package limit

import (
	"sync"
	"time"
)

// DelayLimiter is a rate limiter that will execute a function at most once every X seconds, X being the Delay.
type DelayLimiter struct {
	Delay time.Duration

	mutex      sync.Mutex
	lastCalled time.Time
}

func (l *DelayLimiter) Call(fn func()) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Check if the elapsed time since the last call is less than 10 seconds
	if time.Since(l.lastCalled) < l.Delay {
		return // Discard the call
	}

	// Update the last called time and execute the function
	l.lastCalled = time.Now()
	fn()
}
