package blake3

import (
	"testing"

)

func TestBlake3ConcurrentAggregate(t *testing.T) {
	values := make([]int, 1000); for i := range values { values[i]=1 }
	if got := blake3ConcurrentAggregate(values); got != 1000 { t.Fatalf("sum=%d",got) }
}

func TestBlake3ConcurrentAggregateHandlesEmptyInput(t *testing.T) {
	if got := blake3ConcurrentAggregate(nil); got != 0 { t.Fatalf("empty sum=%d",got) }
}
