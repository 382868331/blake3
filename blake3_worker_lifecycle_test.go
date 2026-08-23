package blake3

import (
	"testing"
	"time"
)

func TestBlake3WorkerLifecycle(t *testing.T) {
	stop:=make(chan struct{}); done:=blake3WorkerLifecycle(stop); close(stop)
	select { case <-done: case <-time.After(50*time.Millisecond): t.Fatal("worker did not stop") }
}

func TestBlake3WorkerLifecycleStopsOnlyAfterSignal(t *testing.T) {
	stop:=make(chan struct{}); done:=blake3WorkerLifecycle(stop)
	select { case <-done: t.Fatal("stopped early"); case <-time.After(2*time.Millisecond): }; close(stop); <-done
}
