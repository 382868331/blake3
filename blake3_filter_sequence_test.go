package blake3

import (
	"testing"
	"reflect"
)

func TestBlake3FilterSequence(t *testing.T) {
	got:=blake3FilterSequence([]int{1,-1,-2,3})
	if !reflect.DeepEqual(got,[]int{1,3}) { t.Fatalf("filtered=%v",got) }
}

func TestBlake3FilterSequenceDoesNotMutateInput(t *testing.T) {
	in:=[]int{1,-1,2}; _=blake3FilterSequence(in)
	if !reflect.DeepEqual(in,[]int{1,-1,2}) { t.Fatalf("input=%v",in) }
}
