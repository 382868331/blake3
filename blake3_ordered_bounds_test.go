package blake3

import (
	"testing"

)

func TestBlake3OrderedBounds(t *testing.T) {
	min,max:=blake3OrderedBounds(9,2)
	if min!=2 || max!=9 { t.Fatalf("range=(%d,%d)",min,max) }
}

func TestBlake3OrderedBoundsKeepsAscendingRange(t *testing.T) {
	min,max:=blake3OrderedBounds(2,9); if min!=2 || max!=9 { t.Fatalf("range=(%d,%d)",min,max) }
}
