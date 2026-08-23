package blake3

import (
	"testing"

)

func TestBlake3ConcurrentAggregate(t *testing.T) {
	values := make([]int, 1000); for i := range values { values[i]=1 }
	if got := blake3ConcurrentAggregate(values); got != 1000 { t.Fatalf("sum=%d",got) }
}
