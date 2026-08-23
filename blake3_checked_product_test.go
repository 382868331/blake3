package blake3

import (
	"testing"
	"math"
)

func TestBlake3CheckedProduct(t *testing.T) {
	if got,err := blake3CheckedProduct(math.MaxInt64/2+1,2); err==nil || got!=0 { t.Fatalf("got=%d err=%v",got,err) }
}
