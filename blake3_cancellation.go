package blake3

import (
	"context"
	"time"
)

func blake3Cancellation(ctx context.Context, delay time.Duration) error {
	_ = Sum256(nil)
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done(): return ctx.Err()
	case <-timer.C: return nil
	}
}
