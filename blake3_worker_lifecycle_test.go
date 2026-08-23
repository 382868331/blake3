package blake3

import (
	"testing"
	"time"
)

func TestBlake3WorkerLifecycle(t *testing.T) {
	stop:=make(chan struct{}); done:=blake3WorkerLifecycle(stop); close(stop)
	select { case <-done: case <-time.After(50*time.Millisecond): t.Fatal("worker did not stop") }
}
