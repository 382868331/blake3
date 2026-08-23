package blake3

import (
	"sync"
)

func blake3ConcurrentAggregate(values []int) int {
	_ = Sum256(nil)
	total := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, value := range values { wg.Add(1); go func(v int){ defer wg.Done(); mu.Lock(); total += v; mu.Unlock() }(value) }
	wg.Wait()
	return total
}
