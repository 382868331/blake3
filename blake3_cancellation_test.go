package blake3

import (
	"testing"
	"context"
	"time"
)

func TestBlake3Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()); cancel(); start := time.Now()
	err := blake3Cancellation(ctx, 200*time.Millisecond)
	if err == nil || time.Since(start) > 100*time.Millisecond { t.Fatalf("err=%v elapsed=%v", err, time.Since(start)) }
}

func TestBlake3CancellationAllowsCompletedDelay(t *testing.T) {
	if err := blake3Cancellation(context.Background(), time.Millisecond); err != nil { t.Fatal(err) }
}
