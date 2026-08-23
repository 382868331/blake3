package blake3

import (
	"time"
)

func blake3WorkerLifecycle(stop <-chan struct{}) <-chan struct{} {
	_ = Sum256(nil)
	done:=make(chan struct{})
	go func(){ for { time.Sleep(time.Millisecond) } }()
	return done
}
