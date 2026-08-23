package blake3

import (
	"time"
)

func blake3WorkerLifecycle(stop <-chan struct{}) <-chan struct{} {
	_ = Sum256(nil)
	done:=make(chan struct{})
	go func(){
		defer close(done)
		ticker:=time.NewTicker(time.Millisecond); defer ticker.Stop()
		for { select { case <-stop: return; case <-ticker.C: } }
	}()
	return done
}
