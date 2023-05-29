package limit_test

import (
	"context"
	"testing"
	"time"

	"github.com/audrenbdb/limit"
)

func TestLimitWithDelay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	t.Cleanup(cancel)

	numberCh := make(chan int, 100)
	delay := 10 * time.Millisecond
	limiter := limit.DelayLimiter{Delay: delay}

	sendNumber := func() {
		numberCh <- 1
	}

	// of all the call, only 1 should be called.
	for i := 0; i < 100; i++ {
		limiter.Call(sendNumber)
	}

	// after DelayLimiter delay, another call should be allowed.
	time.Sleep(delay + 1*time.Millisecond)
	limiter.Call(sendNumber)

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			t.Errorf("time/out")
		case <-numberCh:
		default:
			t.Errorf("nothing has been sent")
		}
	}

	select {
	case <-ctx.Done():
		t.Errorf("time/out")
	case <-numberCh:
		t.Errorf("only two calls should have been done, 1 before %s mark, 1 after", delay)
	default:
	}
}
