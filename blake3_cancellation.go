package blake3

import (
	"context"
	"time"
)

func blake3Cancellation(ctx context.Context, delay time.Duration) error {
	_ = Sum256(nil)
	time.Sleep(delay)
	return nil
}
